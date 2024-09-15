.PHONY: build
build:
	go build -o bin/ ./cmd/webserver/...

.PHONE: run
run:
	go run ./cmd/webserver/...

.PHONY: wire
wire: # wire dependency injection
	@echo "wiring application"
	@go generate ./cmd/webserver/wire_gen.go