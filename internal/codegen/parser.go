package codegen

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type Parameter struct {
	Name      string
	GoName    string
	ParamName string
	Location  string
	Required  bool
	TypeHint  string
	Default   any
	JSONName  string
}

type Endpoint struct {
	Path          string
	Method        string
	OperationID   string
	MethodName    string
	Parameters    []Parameter
	RequestBody   []Parameter
	ResponseModel string
	IsMultipart   bool
}

type Field struct {
	Name     string
	GoName   string
	TypeHint string
	Required bool
	JSONName string
}

type Model struct {
	Name   string
	GoName string
	Fields []Field
}

type OpenAPISpec struct {
	Servers    []struct{ URL string `json:"url"` } `json:"servers"`
	Paths      map[string]map[string]json.RawMessage `json:"paths"`
	Components struct {
		Schemas    map[string]json.RawMessage `json:"schemas"`
		Parameters map[string]json.RawMessage `json:"parameters"`
	} `json:"components"`
}

type ParseResult struct {
	Endpoints []Endpoint
	Models    []Model
	BaseURL   string
}

func ParseSchema(path string) (*ParseResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var spec OpenAPISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	globalSpec = &spec
	result := &ParseResult{}

	if len(spec.Servers) > 0 {
		for _, s := range spec.Servers {
			if strings.Contains(s.URL, "api.") {
				result.BaseURL = s.URL
				break
			}
		}
		if result.BaseURL == "" {
			result.BaseURL = spec.Servers[0].URL
		}
	}

	models, err := parseSchemas(spec)
	if err != nil {
		return nil, err
	}
	result.Models = models

	endpoints, inlineModels, err := parsePaths(spec)
	if err != nil {
		return nil, err
	}
	result.Endpoints = endpoints
	result.Models = append(result.Models, inlineModels...)

	return result, nil
}

func parseSchemas(spec OpenAPISpec) ([]Model, error) {
	var models []Model
	for name, raw := range spec.Components.Schemas {
		var schema map[string]any
		if err := json.Unmarshal(raw, &schema); err != nil {
			return nil, fmt.Errorf("unmarshal schema %s: %w", name, err)
		}

		schemaType, _ := schema["type"].(string)
		_, hasProps := schema["properties"]
		if schemaType != "object" && schemaType != "" && !hasProps {
			continue
		}

		model := Model{
			Name:   name,
			GoName: sanitizeGoName(name),
		}

		props, _ := schema["properties"].(map[string]any)
		requiredList := getStringList(schema, "required")
		requiredSet := toSet(requiredList)

		for propName, propRaw := range props {
			propSchema, _ := propRaw.(map[string]any)
			typeHint := openAPITypeToGo(propSchema)
			isRequired := requiredSet[propName]
			if !isRequired {
				typeHint = makePointer(typeHint)
			}
			model.Fields = append(model.Fields, Field{
				Name:     propName,
				GoName:   toExportedName(propName),
				TypeHint: typeHint,
				Required: isRequired,
				JSONName: propName,
			})
		}
		models = append(models, model)
	}
	return models, nil
}

func parsePaths(spec OpenAPISpec) ([]Endpoint, []Model, error) {
	var endpoints []Endpoint
	var inlineModels []Model
	httpMethods := map[string]bool{"get": true, "post": true, "put": true, "delete": true, "patch": true}

	for path, methods := range spec.Paths {
		for method, raw := range methods {
			if !httpMethods[method] {
				continue
			}

			var opSpec map[string]any
			if err := json.Unmarshal(raw, &opSpec); err != nil {
				return nil, nil, fmt.Errorf("unmarshal operation %s %s: %w", method, path, err)
			}

			operationID, _ := opSpec["operationId"].(string)
			if operationID == "" {
				continue
			}

			ep := Endpoint{
				Path:        path,
				Method:      strings.ToUpper(method),
				OperationID: operationID,
				MethodName:  toPascalCase(operationID),
			}

			if params, ok := opSpec["parameters"].([]any); ok {
				for _, p := range params {
					param, err := resolveParameter(spec, p)
					if err != nil {
						return nil, nil, err
					}
					if param == nil {
						continue
					}
					paramName, _ := (*param)["name"].(string)
					location, _ := (*param)["in"].(string)
					required, _ := (*param)["required"].(bool)
					paramSchema, _ := (*param)["schema"].(map[string]any)
					typeHint := openAPITypeToGo(paramSchema)

					ep.Parameters = append(ep.Parameters, Parameter{
						Name:      paramName,
						GoName:    toExportedName(paramName),
						ParamName: toParamName(paramName),
						Location:  location,
						Required:  required,
						TypeHint:  typeHint,
						JSONName:  paramName,
					})
				}
			}

			if reqBody, ok := opSpec["requestBody"].(map[string]any); ok {
				content, _ := reqBody["content"].(map[string]any)
				for contentType, ctRaw := range content {
					if strings.Contains(contentType, "multipart") {
						ep.IsMultipart = true
					}
					ctSpec, _ := ctRaw.(map[string]any)
					bodySchema, _ := ctSpec["schema"].(map[string]any)
					bodySchema = resolveSchema(spec, bodySchema)

					props, _ := bodySchema["properties"].(map[string]any)
					requiredList := getStringList(bodySchema, "required")
					requiredSet := toSet(requiredList)

					for propName, propRaw := range props {
						propSchema, _ := propRaw.(map[string]any)
						typeHint := openAPITypeToGo(propSchema)
						location := "body"
						format, _ := propSchema["format"].(string)
						if format == "binary" {
							location = "file"
							typeHint = "io.Reader"
						}

						ep.RequestBody = append(ep.RequestBody, Parameter{
							Name:      propName,
							GoName:    toExportedName(propName),
							ParamName: toParamName(propName),
							Location:  location,
							Required:  requiredSet[propName],
							TypeHint:  typeHint,
							JSONName:  propName,
						})
					}
					break
				}
			}

			responses, _ := opSpec["responses"].(map[string]any)
			successResp := getFirstSuccess(responses)
			if successResp != nil {
				content, _ := (*successResp)["content"].(map[string]any)
				for _, ctRaw := range content {
					ctSpec, _ := ctRaw.(map[string]any)
					respSchema, _ := ctSpec["schema"].(map[string]any)
					if ref, ok := respSchema["$ref"].(string); ok {
						parts := strings.Split(ref, "/")
						refName := parts[len(parts)-1]
						ep.ResponseModel = sanitizeGoName(refName)
					} else if respSchema != nil {
						schemaType, _ := respSchema["type"].(string)
						props, hasProps := respSchema["properties"].(map[string]any)
						if schemaType == "object" && hasProps {
							modelName := operationID + "Response"
							goName := sanitizeGoName(modelName)
							model := Model{Name: modelName, GoName: goName}
							requiredList := getStringList(respSchema, "required")
							requiredSet := toSet(requiredList)
							for propName, propRaw := range props {
								propSchema, _ := propRaw.(map[string]any)
								typeHint := openAPITypeToGo(propSchema)
								isReq := requiredSet[propName]
								if !isReq {
									typeHint = makePointer(typeHint)
								}
								model.Fields = append(model.Fields, Field{
									Name:     propName,
									GoName:   toExportedName(propName),
									TypeHint: typeHint,
									Required: isReq,
									JSONName: propName,
								})
							}
							inlineModels = append(inlineModels, model)
							ep.ResponseModel = goName
						}
					}
					break
				}
			}

			endpoints = append(endpoints, ep)
		}
	}

	nameCounts := map[string]int{}
	for _, ep := range endpoints {
		nameCounts[ep.MethodName]++
	}
	nameIdx := map[string]int{}
	for i := range endpoints {
		if nameCounts[endpoints[i].MethodName] > 1 {
			idx := nameIdx[endpoints[i].MethodName]
			nameIdx[endpoints[i].MethodName]++
			suffix := strings.ToUpper(endpoints[i].Method[:1]) + strings.ToLower(endpoints[i].Method[1:])
			if idx > 0 {
				endpoints[i].MethodName = endpoints[i].MethodName + suffix + string(rune('0'+idx))
			} else {
				endpoints[i].MethodName = endpoints[i].MethodName + suffix
			}
		}
	}

	return endpoints, inlineModels, nil
}

func resolveParameter(spec OpenAPISpec, raw any) (*map[string]any, error) {
	m, ok := raw.(map[string]any)
	if !ok {
		return nil, nil
	}
	if ref, ok := m["$ref"].(string); ok {
		parts := strings.Split(ref, "/")
		name := parts[len(parts)-1]
		if paramRaw, ok := spec.Components.Parameters[name]; ok {
			var resolved map[string]any
			if err := json.Unmarshal(paramRaw, &resolved); err != nil {
				return nil, fmt.Errorf("unmarshal parameter %s: %w", name, err)
			}
			return &resolved, nil
		}
		return nil, nil
	}
	return &m, nil
}

func resolveSchema(spec OpenAPISpec, schema map[string]any) map[string]any {
	if ref, ok := schema["$ref"].(string); ok {
		parts := strings.Split(ref, "/")
		name := parts[len(parts)-1]
		if raw, ok := spec.Components.Schemas[name]; ok {
			var resolved map[string]any
			if err := json.Unmarshal(raw, &resolved); err != nil {
				return schema
			}
			return resolved
		}
	}
	return schema
}

func getFirstSuccess(responses map[string]any) *map[string]any {
	for _, code := range []string{"200", "201", "2XX"} {
		if r, ok := responses[code]; ok {
			m, _ := r.(map[string]any)
			return &m
		}
	}
	return nil
}

var globalSpec *OpenAPISpec

func openAPITypeToGo(schema map[string]any) string {
	if schema == nil {
		return "any"
	}
	if ref, ok := schema["$ref"].(string); ok {
		parts := strings.Split(ref, "/")
		refName := parts[len(parts)-1]
		if globalSpec != nil {
			if raw, ok := globalSpec.Components.Schemas[refName]; ok {
				var resolved map[string]any
				if err := json.Unmarshal(raw, &resolved); err != nil {
					return "any"
				}
				schemaType, _ := resolved["type"].(string)
				_, hasProps := resolved["properties"]
				if schemaType != "object" && !hasProps {
					return openAPITypeToGo(resolved)
				}
				return sanitizeGoName(refName)
			}
			return "any"
		}
		return sanitizeGoName(refName)
	}

	schemaType, _ := schema["type"].(string)
	if types, ok := schema["type"].([]any); ok {
		for _, t := range types {
			if s, _ := t.(string); s != "null" {
				schemaType = s
				break
			}
		}
	}

	switch schemaType {
	case "integer":
		return "int64"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "string":
		format, _ := schema["format"].(string)
		if format == "binary" {
			return "io.Reader"
		}
		return "string"
	case "array":
		items, _ := schema["items"].(map[string]any)
		itemType := openAPITypeToGo(items)
		return "[]" + itemType
	case "object":
		return "map[string]any"
	}

	if anyOf, ok := schema["anyOf"].([]any); ok {
		for _, s := range anyOf {
			sm, _ := s.(map[string]any)
			t, _ := sm["type"].(string)
			if t != "null" {
				return openAPITypeToGo(sm)
			}
		}
	}
	if oneOf, ok := schema["oneOf"].([]any); ok {
		for _, s := range oneOf {
			sm, _ := s.(map[string]any)
			t, _ := sm["type"].(string)
			if t != "null" {
				return openAPITypeToGo(sm)
			}
		}
	}

	return "any"
}

func makePointer(t string) string {
	if strings.HasPrefix(t, "[]") || strings.HasPrefix(t, "map[") || t == "any" || t == "io.Reader" {
		return t
	}
	return "*" + t
}

func sanitizeGoName(name string) string {
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, ".", "_")
	return toPascalCase(name)
}

func toPascalCase(s string) string {
	s = regexp.MustCompile(`[^a-zA-Z0-9_]`).ReplaceAllString(s, "_")
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == '.'
	})
	var result strings.Builder
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		result.WriteString(string(runes))
	}
	r := result.String()
	if len(r) == 0 {
		return "X"
	}
	if r[0] >= '0' && r[0] <= '9' {
		return "X" + r
	}
	return r
}

var goReserved = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
	"string": true, "int": true, "bool": true, "error": true, "len": true,
}

func toExportedName(name string) string {
	result := toPascalCase(name)
	lower := strings.ToLower(name)
	if goReserved[lower] {
		result = result + "Value"
	}
	return result
}

func toParamName(name string) string {
	result := strings.ToLower(name[:1]) + toPascalCase(name)[1:]
	if goReserved[result] {
		result = result + "Value"
	}
	return result
}

func getStringList(m map[string]any, key string) []string {
	arr, ok := m[key].([]any)
	if !ok {
		return nil
	}
	var result []string
	for _, v := range arr {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

func toSet(items []string) map[string]bool {
	s := make(map[string]bool)
	for _, item := range items {
		s[item] = true
	}
	return s
}
