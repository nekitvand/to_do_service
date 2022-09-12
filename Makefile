export GO111MODULE=on

.PHONY:	deps
deps:
	go mod download

.PHONY:	build
build:
	CGO_ENABLED=0  go build -o ./bin/grpc-server$(shell go env GOEXE) ./cmd/grpc-server/main.go