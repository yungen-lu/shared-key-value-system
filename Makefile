run:
	go run ./app/main.go
build:
	go build ./app/main.go

doc.swagger:
	swag init -g internal/controller/http/v1/router.go

migrate-up:
	migrate -database 'postgresql://postgres@localhost:5432/postgres?sslmode=disable' -path db/migrations up

migrate-down:
	migrate -database 'postgresql://postgres@localhost:5432/postgres?sslmode=disable' -path db/migrations up

generate:
	go generate ./...

sqlc:
	sqlc generate --experimental