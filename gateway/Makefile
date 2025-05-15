.PHONY: proto
proto:
	rm -f proto/auth/gen/*.go
	protoc \
		--proto_path=../auth/proto \
		--go_out=./proto/auth/gen --go_opt=paths=source_relative \
		--go-grpc_out=./proto/auth/gen --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./proto/auth/gen --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=api/swagger --openapiv2_opt=allow_merge=true,merge_file_name=auth \
		service.proto

.PHONY: server
server:
	go run ./cmd/gateway/

