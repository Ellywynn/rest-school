.PHONY:
.SILENT:

run:
	go run cmd/app/main.go

build:
	go build -o school cmd/app/main.go

db-run:
	docker run --name=online-school -e POSTGRES_PASSWORD='qwerty' -p 5437:5432 -d --rm postgres

migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5437/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5437/postgres?sslmode=disable' down