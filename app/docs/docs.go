// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/users": {
            "get": {
            "tags": [
                "Get Users"
            ],
            "summary": "Find users",
            "description": "Returns multiple users",
            "operationId": "getUsers",
            "parameters": [],
            "responses": {
                "200": {
                "description": "successful operation"
                },
                "400": {
                "description": "Invalid ID supplied"
                },
                "404": {
                "description": "Users not found"
                }
            }
            }
        },
        "/users/{userId}": {
            "get": {
            "tags": [
                "Get User By Id"
            ],
            "summary": "Find user by ID",
            "description": "Returns a single user",
            "operationId": "getUserById",
            "parameters": [
                {
                "name": "userId",
                "in": "path",
                "description": "ID of user to return",
                "required": true,
                "schema": {
                    "type": "integer",
                    "format": "int16"
                }
                }
            ],
            "responses": {
                "200": {
                "description": "successful operation"
                },
                "400": {
                "description": "Invalid ID supplied"
                },
                "404": {
                "description": "User not found"
                }
            }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/v1",
	Schemes:          []string{"http"},
	Title:            "Echo Swagger Example API",
	Description:      "This is a sample server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
