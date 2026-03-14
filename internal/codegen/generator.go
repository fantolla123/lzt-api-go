package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateEndpoint struct {
	Path          string
	Method        string
	MethodName    string
	PathParams    []Parameter
	QueryParams   []Parameter
	BodyParams    []Parameter
	FileParams    []Parameter
	RequiredParams []Parameter
	OptionalParams []Parameter
	IsMultipart   bool
	ResponseModel string
}

type ClientTemplateData struct {
	Package        string
	Endpoints      []TemplateEndpoint
	BaseURL        string
	ResponseModels []string
	HasFileParams  bool
}

type ModelsTemplateData struct {
	Package string
	Models  []Model
}

func Generate(result *ParseResult, outputDir, packageName string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}

	endpoints := prepareEndpoints(result.Endpoints)

	modelSet := map[string]bool{}
	hasFileParams := false
	for _, ep := range endpoints {
		if ep.ResponseModel != "" {
			modelSet[ep.ResponseModel] = true
		}
		if len(ep.FileParams) > 0 {
			hasFileParams = true
		}
	}
	var responseModels []string
	for m := range modelSet {
		responseModels = append(responseModels, m)
	}

	clientData := ClientTemplateData{
		Package:        packageName,
		Endpoints:      endpoints,
		BaseURL:        result.BaseURL,
		ResponseModels: responseModels,
		HasFileParams:  hasFileParams,
	}

	modelsData := ModelsTemplateData{
		Package: packageName,
		Models:  result.Models,
	}

	funcMap := template.FuncMap{
		"lower":    strings.ToLower,
		"toLower":  func(s string) string { return strings.ToLower(s[:1]) + s[1:] },
		"hasPrefix": strings.HasPrefix,
	}

	clientTmpl := template.Must(template.New("client").Funcs(funcMap).Parse(clientTemplate))
	modelsTmpl := template.Must(template.New("models").Funcs(funcMap).Parse(modelsTemplate))

	clientFile, err := os.Create(filepath.Join(outputDir, "client.go"))
	if err != nil {
		return err
	}
	defer clientFile.Close()
	if err := clientTmpl.Execute(clientFile, clientData); err != nil {
		return fmt.Errorf("client template: %w", err)
	}

	modelsFile, err := os.Create(filepath.Join(outputDir, "models.go"))
	if err != nil {
		return err
	}
	defer modelsFile.Close()
	if err := modelsTmpl.Execute(modelsFile, modelsData); err != nil {
		return fmt.Errorf("models template: %w", err)
	}

	fmt.Printf("Generated %d methods in %s\n", len(endpoints), packageName)
	fmt.Printf("Generated %d models\n", len(result.Models))
	fmt.Printf("Output: %s\n", outputDir)

	return nil
}

func prepareEndpoints(endpoints []Endpoint) []TemplateEndpoint {
	var result []TemplateEndpoint

	for _, ep := range endpoints {
		tep := TemplateEndpoint{
			Path:          ep.Path,
			Method:        ep.Method,
			MethodName:    ep.MethodName,
			IsMultipart:   ep.IsMultipart,
			ResponseModel: ep.ResponseModel,
		}

		for _, p := range ep.Parameters {
			switch p.Location {
			case "path":
				tep.PathParams = append(tep.PathParams, p)
			case "query":
				if p.Required {
					tep.RequiredParams = append(tep.RequiredParams, p)
				} else {
					tep.OptionalParams = append(tep.OptionalParams, p)
				}
				tep.QueryParams = append(tep.QueryParams, p)
			}
		}

		for _, p := range ep.RequestBody {
			switch p.Location {
			case "file":
				tep.FileParams = append(tep.FileParams, p)
				if p.Required {
					tep.RequiredParams = append(tep.RequiredParams, p)
				} else {
					tep.OptionalParams = append(tep.OptionalParams, p)
				}
			case "body":
				tep.BodyParams = append(tep.BodyParams, p)
				if p.Required {
					tep.RequiredParams = append(tep.RequiredParams, p)
				} else {
					tep.OptionalParams = append(tep.OptionalParams, p)
				}
			}
		}

		result = append(result, tep)
	}

	return result
}

var clientTemplate = `package {{.Package}}

import (
	"encoding/json"
	"fmt"{{if .HasFileParams}}
	"io"{{end}}
	"net/url"
	"strings"

	lztapi "github.com/fantolla123/lzt-api-go"
)

type Client struct {
	*lztapi.Client
}

func NewClient(token string, opts ...lztapi.Option) *Client {
	return &Client{
		Client: lztapi.NewClient(token, "{{.BaseURL}}", opts...),
	}
}
{{range .Endpoints}}
func (c *Client) {{.MethodName}}({{range .PathParams}}{{.ParamName}} {{.TypeHint}}, {{end}}{{range .RequiredParams}}{{.ParamName}} {{if eq .Location "file"}}io.Reader{{else}}{{.TypeHint}}{{end}}, {{end}}opts ...func(*{{.MethodName}}Params)) ({{if .ResponseModel}}*{{.ResponseModel}}{{else}}map[string]any{{end}}, error) {
	p := &{{.MethodName}}Params{}
	for _, opt := range opts {
		opt(p)
	}

	path := "{{.Path}}"
{{- range .PathParams}}
	path = strings.Replace(path, "{{printf "{%s}" .Name}}", fmt.Sprintf("%v", {{.ParamName}}), 1)
{{- end}}
{{if .QueryParams}}
	params := url.Values{}
{{- range .QueryParams}}
{{- if .Required}}
	params.Set("{{.Name}}", fmt.Sprintf("%v", {{.ParamName}}))
{{- else}}
	if p.{{.GoName}} != nil {
		params.Set("{{.Name}}", fmt.Sprintf("%v", *p.{{.GoName}}))
	}
{{- end}}
{{- end}}
{{- end}}
{{if and .IsMultipart (or .BodyParams .FileParams)}}
	fields := map[string]string{}
{{- range .BodyParams}}
{{- if .Required}}
	fields["{{.Name}}"] = fmt.Sprintf("%v", {{.ParamName}})
{{- else}}
	if p.{{.GoName}} != nil {
		fields["{{.Name}}"] = fmt.Sprintf("%v", *p.{{.GoName}})
	}
{{- end}}
{{- end}}
	files := map[string]io.Reader{}
{{- range .FileParams}}
{{- if .Required}}
	files["{{.Name}}"] = {{.ParamName}}
{{- else}}
	if p.{{.GoName}} != nil {
		files["{{.Name}}"] = p.{{.GoName}}
	}
{{- end}}
{{- end}}

	raw, err := c.DoMultipart("{{.Method}}", path, {{if .QueryParams}}params{{else}}nil{{end}}, fields, files)
{{- else if .BodyParams}}
	body := map[string]any{}
{{- range .BodyParams}}
{{- if .Required}}
	body["{{.Name}}"] = {{.ParamName}}
{{- else}}
	if p.{{.GoName}} != nil {
		body["{{.Name}}"] = *p.{{.GoName}}
	}
{{- end}}
{{- end}}

	raw, err := c.DoRequest("{{.Method}}", path, {{if .QueryParams}}params{{else}}nil{{end}}, body)
{{- else}}
	raw, err := c.DoRequest("{{.Method}}", path, {{if .QueryParams}}params{{else}}nil{{end}}, nil)
{{- end}}
	if err != nil {
		return {{if .ResponseModel}}nil{{else}}nil{{end}}, err
	}
{{if .ResponseModel}}
	data, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	var result {{.ResponseModel}}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
{{- else}}
	return raw, nil
{{- end}}
}

type {{.MethodName}}Params struct {
{{- range .OptionalParams}}
{{- if eq .Location "file"}}
	{{.GoName}} io.Reader
{{- else}}
	{{.GoName}} *{{.TypeHint}}
{{- end}}
{{- end}}
}
{{end}}`

var modelsTemplate = `package {{.Package}}
{{range .Models}}
type {{.GoName}} struct {
{{- range .Fields}}
	{{.GoName}} {{.TypeHint}} ` + "`" + `json:"{{.JSONName}}{{if not .Required}},omitempty{{end}}"` + "`" + `
{{- end}}
}
{{end}}`
