.PHONY: proto
proto:
	rm -f proto/auth/gen/*.go
	protoc \
		--proto_path=../auth/proto \
		--go_out=./proto/auth/gen --go_opt=paths=source_relative \
		--go-grpc_out=./proto/auth/gen --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./proto/auth/gen --grpc-gateway_opt=paths=source_relative \
		service.proto

.PHONY: server
server:
	go run ./cmd/gateway/main.go

