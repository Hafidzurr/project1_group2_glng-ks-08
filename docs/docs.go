// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todos": {
            "get": {
                "description": "Get a list of all Todos",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all Todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Todo"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Todo item",
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new Todo",
                "parameters": [
                    {
                        "description": "Todo object",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get a Todo by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a Todo by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a Todo by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Update a Todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Todo object",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Todo by its ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a Todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
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
        }
    },
    "definitions": {
        "models.Todo": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
