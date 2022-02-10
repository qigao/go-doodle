//go:build tools

package tools

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/grpc-ecosystem/grpc-gateway/utilities"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
