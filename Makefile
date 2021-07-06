code-check: mod imports fmt vet
vet:
	go vet ./...
fmt:
	gofmt -d -s .
imports:
	goimports -w .
mod:
	go mod tidy
	go mod verify
	go mod download


install-oapi-codegen:
	go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen

oapi-codegen:
	oapi-codegen -generate "types" -package gen sushiApi.yaml > ./internal/http/gen/model.go
	oapi-codegen -generate "server,spec" -package gen sushiApi.yaml > ./internal/http/gen/server.go

wire:
	go run github.com/google/wire/cmd/wire ./cmd/...

run-mock:
	go build -ldflags="-w -s" -tags mock -o mock ./cmd
	./mock
run:
	go build -ldflags="-w -s"  -o main ./cmd
	./main
