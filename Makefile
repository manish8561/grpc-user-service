APP_NAME=grpc-user-service

PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/user.proto

SERVER_MAIN=cmd/server/main.go
CLIENT_MAIN=client/main.go

.PHONY: help proto server client tidy fmt test clean

help:
	@echo "Available commands:"
	@echo " make proto   - Generate protobuf and gRPC files"
	@echo " make server  - Run gRPC server"
	@echo " make client  - Run gRPC client"
	@echo " make tidy    - Clean and download dependencies"
	@echo " make fmt     - Format Go code"
	@echo " make test    - Run tests"
	@echo " make clean   - Remove generated protobuf files"

# Generate protobuf files
proto:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_FILE)

# Run gRPC server
server:
	go run $(SERVER_MAIN)

# Run gRPC client
client:
	go run $(CLIENT_MAIN)

# Download and clean dependencies
tidy:
	go mod tidy

# Format all Go files
fmt:
	go fmt ./...

# Run tests
test:
	go test ./...

# Remove generated protobuf files
clean:
	rm -f pb/*.pb.go