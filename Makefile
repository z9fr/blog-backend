TESTNAME ?= $(shell bash -c 'read -p "Test File Name: " testname; echo $$testname')

build:
	go build -o bin/app cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go clean -testcache && go test ./tests -v

# run specific test file using the given name
stest:
	go clean -testcache && go test ./tests/$(TESTNAME) -v

docker:
	sudo docker build -t blogv2-app .
