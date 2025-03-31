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
        "/": {
            "get": {
                "description": "Returns a simple Hello World message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "Home Route",
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
        "/artist": {
            "get": {
                "description": "Gets information for requested artist. If information does not exist in database, it is retrieved from setlist.fm API and entered into database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Artist"
                ],
                "summary": "Serve information for a given artist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Artist Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Object that holds artist information",
                        "schema": {
                            "$ref": "#/definitions/types.Artist"
                        }
                    },
                    "400": {
                        "description": "Error describing failure",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/import": {
            "get": {
                "description": "Gets setlist information from setlist.fm API for given artist, and imports it into database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Artist"
                ],
                "summary": "Import information for a given artist into database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Artist MBID",
                        "name": "mbid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Message indicating success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error describing failure",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserLoginPayload"
                        }
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
                    },
                    "400": {
                        "description": "Invalid email or password",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Register Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserRegisterPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid payload or user already exists",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/validate": {
            "get": {
                "description": "Verifies if a user's session cookie contains an authenticated token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Validate user session",
                "responses": {
                    "200": {
                        "description": "user session validated",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "missing or invalid authorization token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Artist": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mbid": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "types.UserLoginPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.UserRegisterPayload": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 130,
                    "minLength": 3
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Concert Tracker API",
	Description:      "API documentation for Concert Tracker.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
