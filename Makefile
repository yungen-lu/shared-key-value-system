run:
	go run ./app/main.go
build:
	go build ./app/main.go

doc.swagger:
	swag init -g internal/controller/http/v1/router.go

migrate-up:
	migrate -database 'postgresql://postgres@localhost:5432/postgres?sslmode=disable' -path db/migrations up

migrate-down:
	migrate -database 'postgresql://postgres@localhost:5432/postgres?sslmode=disable' -path db/migrations down

generate:
	go generate ./...

sqlc:
	sqlc generate --experimental

test.usecase:
	go test -v ./internal/usecase/...

test.http:
	go test -v ./internal/controller/http/...

test.integration:
	go test -v ./integration-test/integration_test.go

test:
	go test -v ./...

docker.build-image:
	docker build -t yungen/shared-key-value-list-system -f Dockerfile .

critic:
	gocritic check ./...

critic.all:
	gocritic check -enableAll ./...
security:
	gosec ./...

lint:
	golangci-lint run ./...
