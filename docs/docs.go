// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-12-06 22:06:17.730554769 +0800 CST m=+0.149272161

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "An example of gin",
        "title": "Golang Gin API",
        "termsOfService": "https://github.com/Quons/matchmaker",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/Quons/matchmaker/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/articleAndTag": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增文章和标签",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取多个文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":[{\"id\":3,\"created_on\":1516937037,\"modified_on\":0,\"tag_id\":11,\"tag\":{\"id\":11,\"created_on\":1516851591,\"modified_on\":0,\"name\":\"312321\",\"created_by\":\"4555\",\"modified_by\":\"\",\"state\":1},\"content\":\"5555\",\"created_by\":\"2412\",\"modified_by\":\"\",\"state\":1}],\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Desc",
                        "name": "desc",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Content",
                        "name": "content",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "获取单个文章",
                "parameters": [
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"id\":3,\"created_on\":1516937037,\"modified_on\":0,\"tag_id\":11,\"tag\":{\"id\":11,\"created_on\":1516851591,\"modified_on\":0,\"name\":\"312321\",\"created_by\":\"4555\",\"modified_by\":\"\",\"state\":1},\"content\":\"5555\",\"created_by\":\"2412\",\"modified_by\":\"\",\"state\":1},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改文章",
                "parameters": [
                    {},
                    {
                        "type": "string",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Desc",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ModifiedBy",
                        "name": "modified_by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "State",
                        "name": "state",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"请求参数错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除文章",
                "parameters": [
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":400,\"data\":{},\"msg\":\"请求参数错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/getCourse": {
            "post": {
                "description": "获取单个课程description",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "课程"
                ],
                "summary": "获取单个课程",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户token",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "课程ID",
                        "name": "courseId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "10000": {
                        "description": "{\"code\":10000,\"data\":{},\"msg\":\"服务器错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/vo.CourseVo"
                        }
                    },
                    "20000": {
                        "description": "{\"code\":20000,\"data\":{},\"msg\":\"参数错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/import": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "上传图片",
                "parameters": [
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"image_save_url\":\"upload/images/96a.jpg\", \"image_url\": \"http://...\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "vo.CourseVo": {
            "type": "object",
            "properties": {
                "courseId": {
                    "type": "integer"
                },
                "courseImage": {
                    "type": "string"
                },
                "courseName": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
