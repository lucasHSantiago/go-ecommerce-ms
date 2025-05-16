.PHONY: server
server:
	go run ./cmd/product/

.PHONY: grpcui
grpcui:
	grpcui -plaintext localhost:8082

.PHONY: postgres
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -d postgres:alpine

.PHONY: createdb
createdb:
	docker exec -it postgres createdb -U postgres --username=postgres --owner=postgres product-service

.PHONY: dropdb
dropdb:
	docker exec -it postgres dropdb -U postgres product-service

.PHONY: migrateup
migrateup:
	migrate -path internal/infra/migrations -database "postgresql://postgres:admin@localhost:5432/product-service?sslmode=disable" -verbose up $(or $(n))

.PHONY: migratedown
migratedown:
	migrate -path internal/infra/migrations -database "postgresql://postgres:admin@localhost:5432/product-service?sslmode=disable" -verbose down $(or $(n))

.PHONY: new_migration
new_migration:
	migrate create -ext sql -dir internal/infra/migrations -seq $(name)

