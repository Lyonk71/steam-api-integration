run:
	go run .\cmd\main.go

build:
	go build -o bin/main cmd\main.go
	
deps:
	go mod download

lint:
	go fmt ./...
	golangci-lint run