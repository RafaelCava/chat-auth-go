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
            "name": "Rafael Cavalcante",
            "url": "https://github.com/RafaelCava",
            "email": "rafael.software-developer@outlook.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Authenticate a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authenticate a user",
                "parameters": [
                    {
                        "description": "Data to authenticate one user",
                        "name": "AuthRequest",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usecases.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "TokenPair",
                        "schema": {
                            "$ref": "#/definitions/models.TokenPair"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    }
                }
            }
        },
        "/rooms": {
            "post": {
                "description": "Create a new room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Create a new room",
                "parameters": [
                    {
                        "description": "Data to create one room",
                        "name": "CreateRoomParams",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateRoomParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Room",
                        "schema": {
                            "$ref": "#/definitions/models.Room"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Data to create one user",
                        "name": "CreateUserParams",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateUserParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocols.HttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Room": {
            "type": "object",
            "properties": {
                "active_users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.User"
                },
                "owner_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.TokenPair": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "protocols.HttpResponse": {
            "type": "object",
            "properties": {
                "body": {},
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "usecases.AuthRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateRoomParams": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateUserParams": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Authorization header using the Bearer scheme.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Chat Auth Api",
	Description:      "API para gerenciamento de chats, grupos e autenticação de usuários",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
