build:
	@go build -o bin/GO_BACKEND cmd/main.go

test:
	go test -v ./... -short

run: build
	@./bin/GO_BACKEND