go-run:
	go run ./httpd/main.go

go-test:
	go test -v  ./...

go-build:
	go mod tidy && go build -o bin/ ./httpd/main.go

docker-build:
	docker build -t ginapp:1 .;

docker-run:
	docker run --name=ginapp --rm -p 8080:8080 ginapp:1
