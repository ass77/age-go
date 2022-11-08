.PHONY: run tidy

run:
	go run cmd/app/main.go

build: 
	go build -o cmd/app/grapher cmd/app/main.go

run-build:
	./cmd/app/grapher

tidy:
	go mod tidy

test:
	cd age && go test -v 