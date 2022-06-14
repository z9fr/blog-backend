build:
	go build -o bin/main cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go clean -testcache && go test ./tests -v
