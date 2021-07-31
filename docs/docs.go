// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/city": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Show a city list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search city by str",
                        "name": "str",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "search city by country",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "query limit",
                        "name": "lim",
                        "in": "query"
                    }
                ]
            }
        },
        "/country": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Show a country list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search country by str",
                        "name": "str",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "query limit",
                        "name": "lim",
                        "in": "query"
                    }
                ]
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "188.225.72.165",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "Swagger Geo API",
	Description: "This is a api server for geo clue.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
