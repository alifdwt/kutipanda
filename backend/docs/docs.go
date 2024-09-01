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
            "name": "Alif Dewantara",
            "url": "https://github.com/alifdwt",
            "email": "aputradewantara@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/hello": {
            "get": {
                "description": "Return a greeting message to the user",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Greet the user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login to Kutipanda",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register to Kutipanda",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Full Name",
                        "name": "full_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Description",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Profile Image",
                        "name": "profile_image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/country": {
            "get": {
                "description": "Retrieve all countries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "Get all countries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/country/code/{code}": {
            "get": {
                "description": "Retrieve a country by its code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "Get country by code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Country Code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/country/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "Create country",
                "parameters": [
                    {
                        "description": "Request body to create a new country",
                        "name": "country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/country.CreateCountryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/country/hello": {
            "get": {
                "description": "Return a greeting message for country",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "Greet the user for country",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/movie": {
            "get": {
                "description": "Retrieve all movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get all movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/movie/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Create movie",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Movie description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Movie year",
                        "name": "year",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Movie poster image",
                        "name": "poster_image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Movie country id",
                        "name": "country_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/movie/hello": {
            "get": {
                "description": "Return a greeting message for movie",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Greet the user for movie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/movie/slug/{slug}": {
            "get": {
                "description": "Retrieve a movie by its slug",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get movie by slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/movie/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve movie data by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get movie by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/quote": {
            "get": {
                "description": "Get all quotes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Get all quotes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/quote/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create quote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Create quote",
                "parameters": [
                    {
                        "description": "Create quote",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/quote.CreateQuoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/quote/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete quote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Delete quote",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quote ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/quote/hello": {
            "get": {
                "description": "Greet hello",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Greet hello",
                "responses": {
                    "200": {
                        "description": "Hello from kutipanda",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song": {
            "get": {
                "description": "Get all songs from application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Get all songs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page ID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order By",
                        "name": "order_by",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/song/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new song for the authenticated user",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Create Song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Song lyrics",
                        "name": "lyrics",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Song year",
                        "name": "year",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Song album image",
                        "name": "album_image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Song country id",
                        "name": "country_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/song/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a song by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Delete song by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/song/slug/{slug}": {
            "get": {
                "description": "Retrieve a song by its slug",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Get Song by Slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song Slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve logged in user data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/user/username/{username}": {
            "get": {
                "description": "Retrieve user data by username",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user data by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginRequest": {
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
        "country.CreateCountryRequest": {
            "type": "object",
            "required": [
                "code",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 2
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "quote.CreateQuoteRequest": {
            "type": "object",
            "required": [
                "movie_id",
                "quote_text"
            ],
            "properties": {
                "movie_id": {
                    "type": "integer"
                },
                "quote_text": {
                    "type": "string"
                }
            }
        },
        "responses.ErrorMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "responses.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "Header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Kutipanda API",
	Description:      "REST API for Kutipanda App",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
