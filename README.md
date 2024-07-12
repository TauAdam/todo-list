# todo-list
## rest api for todo list

### Installation
1. Clone the repository
2. Run `go mod tidy`
3. Run docker compose `docker-compose up -d`
4. Run `go run cmd/main.go`

## Features
- [x] Create a todo list
- [] Get all todo list
- [] Get a todo list
- [] Update a todo list
- [] Delete a todo list
- [] Create a todo item
- [] Get all todo items
- [] Get a todo item
- [] Update a todo item
- [] Delete a todo item
- [] Mark a todo item as done
- [x] Sign up
- [x] Sign in

## Tech Stack
- Go
  - Gin
  - Sqlx
  - JWT
  - Viper
  - godotenv
- Postgres