postgres:
	docker run --name small_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine
createdb:
	docker exec -it small_bank createdb --username=root --owner=root small_bank
dropdb:
	docker exec -it small_bank dropdb small_bank
migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/small_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/small_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

format:
	go fmt ./...





.PHONY: postgres createdb dropdb migrateup migratedown sqlc test