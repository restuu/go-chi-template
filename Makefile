.PHONY: build
build:
	go build -o bin/ ./cmd/webserver/...

.PHONE: run
run:
	go run ./cmd/webserver/...