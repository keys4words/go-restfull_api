
build:
	go build -v ./cmd/api

migrate_up:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up

migrate_down:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down

.DEFAULT_GOAL:=build