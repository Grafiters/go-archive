// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "This is the root endpoint.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Welcome to Buffalo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/public": {
            "get": {
                "description": "This is the hello endpoint.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public"
                ],
                "summary": "Say hello to Buffalo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ordering item by request (default: id)",
                        "name": "orderby",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sorting item by ascending or descending (default: asc)",
                        "name": "sortby",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit is the maximum number of items per page (default: 10)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Number of items in page (default: 1)",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.PaginationParam"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "interfaces.PaginationParam": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "Limit is the maximum number of items per page (default: 10)",
                    "type": "integer"
                },
                "page": {
                    "description": "Page Number of items in page (default: 1)",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
