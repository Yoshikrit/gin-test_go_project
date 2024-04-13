// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Walter White",
            "url": "https://twitter.com/example",
            "email": "example@gmail.com"
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
        "/producttype/": {
            "get": {
                "description": "Get ProductTypes from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Get ProductTypes",
                "responses": {
                    "200": {
                        "description": "Get ProductTypes Successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create ProductType to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Create ProductType",
                "responses": {
                    "200": {
                        "description": "Create ProductType Successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/producttype/count": {
            "get": {
                "description": "Get ProductType's count from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Get ProductType's Count",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/producttype/health": {
            "get": {
                "description": "Check if the service is up and running",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Check health of the service",
                "responses": {
                    "200": {
                        "description": "ProductType Service : OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/producttype/{id}": {
            "get": {
                "description": "Get ProductType from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Get ProductType",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get ProductType Successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update ProductType from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Update ProductType",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update ProductType Successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete ProductType from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Delete ProductType",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete ProductType Successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Gin-Test API",
	Description:      "Gin-Test API for testing.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
