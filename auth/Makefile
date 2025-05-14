.PHONY: server
server:
	go run ./cmd/auth/main.go

.PHONY: grpcui
grpcui:
	grpcui -plaintext localhost:8080

.PHONY: postgres
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -d postgres:alpine

.PHONY: createdb
createdb:
	docker exec -it postgres createdb -U postgres --username=postgres --owner=postgres auth-service

.PHONY: dropdb
dropdb:
	docker exec -it postgres dropdb -U postgres auth-service

.PHONY: migrateup
migrateup:
	migrate -path internal/infra/migrations -database "postgresql://postgres:admin@localhost:5432/auth-service?sslmode=disable" -verbose up $(or $(n))

.PHONY: migratedown
migratedown:
	migrate -path internal/infra/migrations -database "postgresql://postgres:admin@localhost:5432/auth-service?sslmode=disable" -verbose down $(or $(n))

.PHONY: new_migration
new_migration:
	migrate create -ext sql -dir internal/infra/migrations -seq $(name)

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: mock
mock:
	mockgen -package application -destination internal/application/mock/distributor.go github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application TaskDistributor
	mockgen -package application -destination internal/application/mock/user_repository.go github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application UserRepository
	mockgen -package application -destination internal/application/mock/session_repository.go github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application SessionRepository
	mockgen -package application -destination internal/application/mock/verify_email_repository.go github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application VerifyEmailRepository


.PHONY: redis
redis:
	docker run --name redis -p 6379:6379 -d redis:7.4.2-alpine

.PHONY: proto
proto:
	rm -f proto/gen/*.go
	protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: test
test:
	go test -v -cover -short ./...

