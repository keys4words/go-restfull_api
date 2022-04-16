run:
	docker-compose up -d
	sleep 20
	go run ./cmd/main.go

lint:
	@golangci-lint run ./... -v

lint_autofix:
	@GO111MODULE=on $(GOLINT) run ./... -v --fix