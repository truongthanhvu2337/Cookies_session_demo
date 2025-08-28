all: build test

build:
	@echo "Building..."
	
	
	@go build -o main.exe cmd/main.go

run:
	@go run cmd/api/main.go