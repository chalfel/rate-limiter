test:
	go test ./... -cover

build: test
	go build -o ./bin/rate-limiter

run: build
	./bin/rate-limiter serve