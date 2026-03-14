package lztapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	maxRetries  = 5
	backoffBase = 1.0
	backoffMax  = 60.0
)

var retryStatuses = map[int]bool{429: true, 502: true, 503: true}

type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

type Option func(*Client)

func WithProxy(proxyURL string) Option {
	return func(c *Client) {
		u, err := url.Parse(proxyURL)
		if err != nil {
			panic(fmt.Sprintf("lztapi: invalid proxy URL: %v", err))
		}
		transport := &http.Transport{Proxy: http.ProxyURL(u)}
		c.httpClient.Transport = transport
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

func NewClient(token, baseURL string, opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    strings.TrimRight(baseURL, "/"),
		token:      token,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) DoRequest(method, path string, params url.Values, body io.Reader) ([]byte, error) {
	fullURL := c.baseURL + path
	if params != nil && len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = io.ReadAll(body)
		if err != nil {
			return nil, err
		}
	}

	for attempt := range maxRetries {
		var reqBody io.Reader
		if bodyBytes != nil {
			reqBody = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequest(method, fullURL, reqBody)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
		if bodyBytes != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		if retryStatuses[resp.StatusCode] {
			if attempt == maxRetries-1 {
				if resp.StatusCode == 429 {
					return nil, &RateLimitError{RetryAfter: getRetryDelay(resp, attempt)}
				}
				return nil, handleErrorBytes(resp.StatusCode, respBody)
			}
			delay := getRetryDelay(resp, attempt)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			continue
		}

		if err := handleErrorBytes(resp.StatusCode, respBody); err != nil {
			return nil, err
		}

		return respBody, nil
	}
	return nil, fmt.Errorf("lztapi: max retries exceeded")
}

func (c *Client) DoMultipart(method, path string, params url.Values, fields map[string]string, files map[string]io.Reader) ([]byte, error) {
	fullURL := c.baseURL + path
	if params != nil && len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for k, v := range fields {
		writer.WriteField(k, v)
	}
	for name, reader := range files {
		part, err := writer.CreateFormFile(name, name)
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, reader); err != nil {
			return nil, err
		}
	}
	writer.Close()

	bodyBytes := buf.Bytes()
	contentType := writer.FormDataContentType()

	for attempt := range maxRetries {
		req, err := http.NewRequest(method, fullURL, bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Content-Type", contentType)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		if retryStatuses[resp.StatusCode] {
			if attempt == maxRetries-1 {
				if resp.StatusCode == 429 {
					return nil, &RateLimitError{RetryAfter: getRetryDelay(resp, attempt)}
				}
				return nil, handleErrorBytes(resp.StatusCode, respBody)
			}
			delay := getRetryDelay(resp, attempt)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			continue
		}

		if err := handleErrorBytes(resp.StatusCode, respBody); err != nil {
			return nil, err
		}

		return respBody, nil
	}
	return nil, fmt.Errorf("lztapi: max retries exceeded")
}

func handleErrorBytes(statusCode int, body []byte) error {
	if statusCode >= 200 && statusCode < 300 {
		return nil
	}
	if statusCode == 401 {
		return &AuthError{}
	}
	if statusCode == 404 {
		return &NotFoundError{}
	}
	if statusCode >= 400 {
		msg := string(body)
		var errResp struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Error.Message != "" {
			msg = errResp.Error.Message
		}
		return &APIError{StatusCode: statusCode, Message: msg}
	}
	return nil
}

func getRetryDelay(resp *http.Response, attempt int) float64 {
	if ra := resp.Header.Get("Retry-After"); ra != "" {
		if v, err := strconv.ParseFloat(ra, 64); err == nil {
			return v
		}
	}
	delay := backoffBase * math.Pow(2, float64(attempt))
	if delay > backoffMax {
		delay = backoffMax
	}
	return delay
}
