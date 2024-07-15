## Implemented rest api for todo list app with:
- Postgres DB
- SQL queries
- Clean architecture 
- JWT for authentication
- Swagger documentation

### Installation
1. Clone the repository
2. Run `go mod tidy`
3. Run docker compose `docker-compose up -d`
4. Run `go run cmd/main.go`

## Features
- [x] Create a todo list
- [x] Get all todo list
- [x] Get a todo list
- [x] Update a todo list
- [x] Delete a todo list
- [x] Create a todo item
- [x] Get all todo items
- [x] Get a todo item
- [x] Update a todo item
- [x] Delete a todo item
- [ ] Mark a todo item as done
- [x] Sign up
- [x] Sign in

## Tech Stack
- Go
  - Gin
  - Sqlx
  - JWT
  - Viper
  - godotenv
- Docker <- Postgres