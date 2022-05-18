# franchise-ordering-system

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

# Table of Content
- [Description](#description)
- [How to Use](#how-to-use)
- [Database Schema](#database-schema)
- [Testing Coverage]($testing-coverage)
- [Feature](#feature)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
Application for managing queue of order in franchise, restaurant, café, etc.

# Database Schema

# Testing Coverage

# Feature

# How to Use
- Installing dependencies
```bash
go mod tidy
```
- Run Local
```bash
go run .\cmd\api\main.go
```
- swagger-doc: generating swagger
```bash
$ swag init --generalInfo cmd/api/main.go --dir ./ --output docs/swagger --exclude logs,web,assets,database
```
- swagger-fmt: for formatting swagger in controllers
```bash
$ swag fmt --generalInfo cmd/api/main.go --dir ./ --exclude logs,web,assets,database
```

# Endpoints

# Credits
- [Muchammad Abdurrochman](https://github.com/Abdurrochman25) (Author)
- [Fatih Al-Fikri](https://github.com/afaf-tech) (Author)
