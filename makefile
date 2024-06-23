run:
	@go run ./cmd/api

test:
	@GOFLAGS="-COUNT=1" go test -v -cover -race ./...
