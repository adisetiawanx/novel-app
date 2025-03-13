# Novel App

This is an application server built using the Go programming language. This app can perform CRUD operations and various other features.

## Tech Stack
- Echo (https://github.com/labstack/echo)
- Validator (https://github.com/go-playground/validator)
- Gorm (https://github.com/go-gorm/gorm)
- Viper (https://github.com/spf13/viper)

## Installation Guide

### 1. Set Up Environment Variables
Create a `.env` file in the `/internal/app/config` directory by copying the example file from the `/internal/app/config/.env.example`

### 2. Set Up Database Migration
coming soon..

### 3. Run the Application

Ensure your database connection is properly configured before starting the server.

Run the following command:

```sh
go run ./cmd/server
```
