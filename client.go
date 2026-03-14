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
			return
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

func (c *Client) DoRequest(method, path string, params url.Values, body any) (map[string]any, error) {
	fullURL := c.baseURL + path
	if params != nil && len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(data)
	}

	for attempt := range maxRetries {
		req, err := http.NewRequest(method, fullURL, bodyReader)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if retryStatuses[resp.StatusCode] {
			if attempt == maxRetries-1 {
				if resp.StatusCode == 429 {
					return nil, &RateLimitError{}
				}
				return nil, handleError(resp)
			}
			delay := getRetryDelay(resp, attempt)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			if bodyReader != nil {
				if seeker, ok := bodyReader.(io.Seeker); ok {
					seeker.Seek(0, io.SeekStart)
				}
			}
			continue
		}

		if err := handleError(resp); err != nil {
			return nil, err
		}

		var result map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, fmt.Errorf("lztapi: max retries exceeded")
}

func (c *Client) DoMultipart(method, path string, params url.Values, fields map[string]string, files map[string]io.Reader) (map[string]any, error) {
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

	for attempt := range maxRetries {
		req, err := http.NewRequest(method, fullURL, bytes.NewReader(buf.Bytes()))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if retryStatuses[resp.StatusCode] {
			if attempt == maxRetries-1 {
				if resp.StatusCode == 429 {
					return nil, &RateLimitError{}
				}
				return nil, handleError(resp)
			}
			delay := getRetryDelay(resp, attempt)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			continue
		}

		if err := handleError(resp); err != nil {
			return nil, err
		}

		var result map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, fmt.Errorf("lztapi: max retries exceeded")
}

func handleError(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	if resp.StatusCode == 401 {
		return &AuthError{}
	}
	if resp.StatusCode == 404 {
		return &NotFoundError{}
	}
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		var errResp struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		msg := string(body)
		if json.Unmarshal(body, &errResp) == nil && errResp.Error.Message != "" {
			msg = errResp.Error.Message
		}
		return &APIError{StatusCode: resp.StatusCode, Message: msg}
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
