create-migration:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema/ -database 'postgres://postgres:{$$DATABASE_PASSWORD}@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema/ -database 'postgres://postgres:{$$DATABASE_PASSWORD}@localhost:5432/postgres?sslmode=disable' down

test:
	go test -v ./... -coverprofile cover.out
test-coverage:
	go tool cover -func cover.out | grep total | awk '{print $3}'