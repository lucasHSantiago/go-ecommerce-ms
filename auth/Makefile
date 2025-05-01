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
	mockgen -source internal/application/user_repository.go -destination internal/application/mock/user_repository.go -package mockdb
	mockgen -source internal/application/session_repository.go -destination internal/application/mock/session_repository.go -package mockdb
	mockgen -source internal/application/distributor/distributor.go -destination internal/application/distributor/mock/distributor.go -package mockwk

.PHONY: redis
redis:
	docker run --name redis -p 6379:6379 -d redis:7.4.2-alpine

#@TODO: need to change this
.PHONY: proto
proto:
	rm -f proto/gen/*.go
	protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: test
test:
	go test -v -cover -short ./...

