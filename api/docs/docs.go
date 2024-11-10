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
        "/admin/user": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Error creating user",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logs in a user and generates a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token and role",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Could not generate token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/servers": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Retrieve a list of all servers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "List all servers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Server"
                            }
                        }
                    },
                    "500": {
                        "description": "Error fetching servers",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/servers/add": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create a new server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Add a new server",
                "parameters": [
                    {
                        "description": "Server details",
                        "name": "server",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Server"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Server created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Error creating a server",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/servers/request-access": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Submits an access request for a server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Request access to a server",
                "parameters": [
                    {
                        "description": "Server access request details",
                        "name": "serverRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AccessRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Access request submitted",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Error requesting access",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AccessRequestDTO": {
            "type": "object",
            "properties": {
                "server_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Server": {
            "type": "object",
            "properties": {
                "environment": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "description": "Role can be \"admin\" or \"employee\"",
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "EnvironmentManager",
	Description:      "Api for Environment access management.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}