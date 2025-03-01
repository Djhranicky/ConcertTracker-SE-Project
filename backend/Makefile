build:
	@go build -o bin/GO_BACKEND cmd/main.go

test:
	@go test github.com/djhranicky/ConcertTracker-SE-Project/routes
	@go test github.com/djhranicky/ConcertTracker-SE-Project/service/auth

run: build
	@./bin/GO_BACKEND