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
        "/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回包括用户信息,token,过期时间",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.Login"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.Login": {
            "type": "object",
            "properties": {
                "captcha": {
                    "description": "验证码",
                    "type": "string"
                },
                "captchaId": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "secret": {
                    "description": "由用户名和密码组成，例如: {\"username\":\"oms_admin\",\"password\":\"Oms123Admin456\"}",
                    "type": "string"
                }
            }
        },
        "response.Login": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/response.LoginUser"
                }
            }
        },
        "response.LoginGroups": {
            "type": "object",
            "properties": {
                "orgCode": {
                    "type": "string"
                },
                "shortName": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "sysRoles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.LoginRole"
                    }
                }
            }
        },
        "response.LoginRole": {
            "type": "object",
            "properties": {
                "roleCode": {
                    "type": "string"
                },
                "roleName": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "response.LoginUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "logOperation": {
                    "type": "boolean"
                },
                "nickName": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "sysGroups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.LoginGroups"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        },
        "ApiKeyDomain": {
            "type": "apiKey",
            "name": "x-group",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "V1.0.0",
	Host:             "127.0.0.1:8888",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Oms-Admin Swagger API接口文档",
	Description:      "使用gin的全栈开发基础平台",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
