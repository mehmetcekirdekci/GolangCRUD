// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/product": {
            "put": {
                "tags": [
                    "product"
                ],
                "summary": "Updates product by Id",
                "parameters": [
                    {
                        "description": "UpdateProductRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "product"
                ],
                "summary": "Creates product",
                "parameters": [
                    {
                        "description": "CreateProductRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.BaseResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.CreateProductRequest": {
            "type": "object",
            "required": [
                "name",
                "price"
            ],
            "properties": {
                "currency": {
                    "maximum": 2,
                    "minimum": 0,
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.CurrencyTypeEnum"
                        }
                    ]
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "controller.UpdateProductRequest": {
            "type": "object",
            "required": [
                "id",
                "name",
                "price"
            ],
            "properties": {
                "currency": {
                    "minimum": 0,
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.CurrencyTypeEnum"
                        }
                    ]
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "types.CurrencyTypeEnum": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "Dollar",
                "Euro",
                "Pound"
            ]
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
