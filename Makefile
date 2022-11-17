.PHONY: run tidy

run:
	go run cmd/app/main.go

local-build: 
	go build -o cmd/app/grapher cmd/app/main.go

local-run-build:
	./cmd/app/grapher

tidy:
	go mod tidy

test:
	cd age && go test -v 

docker-build-local:
	docker build --build-arg DSN="host=127.0.0.1 port=5432 dbname=test user=postgres password=postgres sslmode=disable" \
	-t grapher-local -f Dockerfile.local .

docker-run-local:
	docker run -p 5000:5000 grapher-local:latest