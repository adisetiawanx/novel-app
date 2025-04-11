# Novel App

This is an application server built using the Go programming language. This app can perform CRUD operations and various other features.

## External Libraries
- Echo (https://github.com/labstack/echo)
- Validator (https://github.com/go-playground/validator)
- Gorm (https://github.com/go-gorm/gorm)
- Gorm Postgres Driver (https://github.com/go-gorm/postgres)
- Viper (https://github.com/spf13/viper)
- Golang JWT (https://github.com/golang-jwt/jwt)
- Google UUID (https://github.com/google/uuid)
- Bcrypt (https://golang.org/x/crypto/bcrypt)


## Installation Guide

### 1. Set Up Environment Variables
Create a `.env` file in the `/configs` directory by copying the example file from the `/configs/.env.example`

### 2. Set Up Database Migration

#### 1. Install Golang Migrate
You need `golang-migrate` to run database migrations. 

#### Using Homebrew (MacOS/Linux)
```sh
brew install golang-migrate
```

#### Using Chocolatey (Windows)
```sh
choco install golang-migrate
```

#### Using Go
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

#### Set Up Environment Variables (Windows)
If you installed `golang-migrate` using Go (`go install`), you need to add it to your system's **environment variables** to use the `migrate` command globally.

1. Open **System Properties** → **Advanced** → **Environment Variables**.
2. Under **System variables**, find `Path` and click **Edit**.
3. Click **New** and add:
   ```
   C:\Users\YourUsername\go\bin
   ```
4. Click **OK**, then **restart your terminal**.
5. Verify the installation:
   ```sh
   migrate -version
   ```

For more details, visit the official **golang-migrate** repository (https://github.com/golang-migrate/migrate)

#### 2. Apply the Migrations
```sh
migrate -path migrations -database "postgres://<YOUR_DB_USER>:<YOUR_DB_PASSWORD>@<YOUR_DB_HOST>:<YOUR_DB_PORT>/<YOUR_DB_NAME>?sslmode=disable" -verbose up
```
example
```sh
migrate -path migrations -database "postgres://root:dev@localhost:5432/novel_app?sslmode=disable" -verbose up
```
#### Rollback Migrations
```sh
migrate -path migrations -database "postgres://<YOUR_DB_USER>:<YOUR_DB_PASSWORD>@<YOUR_DB_HOST>:<YOUR_DB_PORT>/<YOUR_DB_NAME>?sslmode=disable" -verbose down
```
example
```sh
migrate -path migrations -database "postgres://root:dev@localhost:5432/novel_app?sslmode=disable" -verbose down
```

### 3. Run the Application

Ensure your database connection is properly configured before starting the server.

Run the following command:

```sh
go run ./cmd/novel-app
```

