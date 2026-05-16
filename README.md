````markdown
# gRPC User Service in Golang

A production-style gRPC microservice example built using Golang.

This project demonstrates:

- gRPC server setup
- protobuf integration
- clean architecture
- dependency injection
- repository pattern
- service layer separation
- unary RPC implementation

---

# Project Structure

```txt
grpc-user-service/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── client/
│   └── main.go
│
├── internal/
│   ├── handler/
│   │   └── grpc/
│   │       └── user_handler.go
│   │
│   ├── model/
│   │   └── user.go
│   │
│   ├── repository/
│   │   └── user_repository.go
│   │
│   └── service/
│       └── user_service.go
│
├── pb/
│   ├── user.pb.go
│   └── user_grpc.pb.go
│
├── proto/
│   └── user.proto
│
├── pkg/
│   └── logger/
│       └── logger.go
│
├── go.mod
├── Makefile
└── README.md
````

---

# Requirements

Install the following:

* Golang 1.22+
* Protocol Buffers Compiler (`protoc`)
* gRPC Go plugins

---

# Install protoc

## Ubuntu

```bash
sudo apt install -y protobuf-compiler
```

## Mac

```bash
brew install protobuf
```

Verify installation:

```bash
protoc --version
```

---

# Install Go gRPC Plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Add Go binaries to PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

# Initialize Project

```bash
go mod init grpc-user-service
```

Install dependencies:

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

---

# Proto Definition

File:

```txt
proto/user.proto
```

```proto
syntax = "proto3";

package user;

option go_package = "./pb";

service UserService {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
}

message GetUserRequest {
  int64 id = 1;
}

message GetUserResponse {
  int64 id = 1;
  string name = 2;
  string email = 3;
}
```

---

# Generate gRPC Files

Run:

```bash
protoc \
  --go_out=. \
  --go-grpc_out=. \
  proto/user.proto
```

Generated files:

```txt
pb/user.pb.go
pb/user_grpc.pb.go
```

---

# Run gRPC Server

```bash
go run cmd/server/main.go
```

Expected output:

```txt
gRPC server running on :50051
```

---

# Run gRPC Client

Open another terminal:

```bash
go run client/main.go
```

Expected response:

```txt
id:1 name:"Manish" email:"manish@test.com"
```

---

# Architecture Overview

```txt
Client
   ↓
gRPC Handler
   ↓
Service Layer
   ↓
Repository Layer
   ↓
Database
```

---

# Layer Responsibilities

## Handler Layer

Location:

```txt
internal/handler/grpc/
```

Responsibilities:

* gRPC transport handling
* request parsing
* response formatting
* status/error mapping

---

## Service Layer

Location:

```txt
internal/service/
```

Responsibilities:

* business logic
* validations
* orchestration
* workflow handling

---

## Repository Layer

Location:

```txt
internal/repository/
```

Responsibilities:

* database queries
* persistence logic
* external storage interaction

---

# Why Use gRPC?

Advantages:

* high performance
* binary protocol (protobuf)
* smaller payloads
* HTTP/2 support
* streaming support
* strongly typed contracts
* code generation

---

# gRPC Communication Types

| Type                    | Description                         |
| ----------------------- | ----------------------------------- |
| Unary                   | Single request → Single response    |
| Server Streaming        | Single request → Multiple responses |
| Client Streaming        | Multiple requests → Single response |
| Bidirectional Streaming | Both sides stream                   |

---

# Production Enhancements

In enterprise systems, you should additionally implement:

* authentication
* authorization
* TLS encryption
* structured logging
* tracing
* metrics
* retries
* rate limiting
* graceful shutdown
* service discovery
* health checks
* circuit breakers

---

# Example Unary Interceptor

```go
grpc.NewServer(
    grpc.UnaryInterceptor(loggingInterceptor),
)
```

Common interceptors:

* logging
* authentication
* tracing
* metrics
* panic recovery

---

# Example Makefile

```makefile
proto:
	protoc \
	--go_out=. \
	--go-grpc_out=. \
	proto/user.proto

server:
	go run cmd/server/main.go

client:
	go run client/main.go
```

---

# Example Commands

Generate protobuf:

```bash
make proto
```

Run server:

```bash
make server
```

Run client:

```bash
make client
```

---

# Common Production Folder Additions

```txt
configs/
deployments/
scripts/
migrations/
tests/
docs/
```

---

# Example Real-World Stack

```txt
API Gateway
    ↓
Auth Service
    ↓
User Service (gRPC)
    ↓
PostgreSQL

Order Service
    ↓
Kafka
    ↓
Notification Service
```

---

# Technologies Used

* Golang
* gRPC
* Protocol Buffers
* HTTP/2

---

# Useful Commands

Format code:

```bash
go fmt ./...
```

Run tests:

```bash
go test ./...
```

Download dependencies:

```bash
go mod tidy
```

---

# Future Improvements

* Add PostgreSQL
* Add Redis caching
* Add JWT authentication
* Add Docker support
* Add Kubernetes deployment
* Add OpenTelemetry tracing
* Add Prometheus metrics
* Add CI/CD pipeline

---

# License

MIT

```
```
