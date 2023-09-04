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
        "/conf/get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "es配置"
                ],
                "summary": "获取es配置",
                "responses": {
                    "200": {
                        "description": "code\",\"msg\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/conf/set": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "es配置"
                ],
                "summary": "设置es配置",
                "parameters": [
                    {
                        "description": "EsConfig",
                        "name": "esConf",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EsConfig"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"msg\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/conf/use": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "es配置"
                ],
                "summary": "应用es配置",
                "parameters": [
                    {
                        "description": "EsConfig",
                        "name": "esConf",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EsConfig"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"msg\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/es/getMapping": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "es查询"
                ],
                "summary": "获取索引字段",
                "responses": {
                    "200": {
                        "description": "code\",\"msg\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/index": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "首页"
                ],
                "summary": "首页",
                "responses": {
                    "200": {
                        "description": "code\",\"msg\",\"data\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.EsConfig": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
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
