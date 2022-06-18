// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Login with Email and Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Email: User Email; Password: User Password",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Auth"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorBadRequestValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Register with Email and Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Email: User Email; Password: User Password",
                        "name": "registerReqest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorBadRequestValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/customers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create New Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Create New Customer",
                "parameters": [
                    {
                        "description": "Name: Customer name; Email: Customer email; Phone: Customer phone number",
                        "name": "customerRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateCustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Customer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorBadRequestValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/foods": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get All Food With Pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Food"
                ],
                "summary": "Get All Food With Pagination",
                "parameters": [
                    {
                        "description": "Page: Page Number; Limit: Request Limit; Search: Reqeust for search food",
                        "name": "foodRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetAllFoodRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorBadRequestValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/HttpErrorInternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "APIResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "page_info": {}
            }
        },
        "HttpErrorBadRequestValidation": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "field": {
                                "type": "string",
                                "example": "email"
                            },
                            "message": {
                                "type": "string",
                                "example": "The email must be a valid email address."
                            }
                        }
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Error Validation"
                }
            }
        },
        "HttpErrorInternalServerError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "HttpErrorNotFound": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "record not found"
                }
            }
        },
        "HttpErrorUnauthorized": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "missing key in request header"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Auth": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Customer": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "request.AuthRequest": {
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
        "request.CreateCustomerRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "request.GetAllCustomerRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "This is limit, if load all data set value to -1",
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "description": "This is page, if load all data set value to 0",
                    "type": "integer",
                    "example": 1
                },
                "search": {
                    "type": "string"
                }
            }
        },
        "request.GetAllFoodRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "This is limit, if load all data set value to -1",
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "description": "This is page, if load all data set value to 0",
                    "type": "integer",
                    "example": 1
                },
                "search": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "franchise-ordering-system.swagger.io",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Franchise Ordering System API",
	Description:      "This is a franchise-ordering-system server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
