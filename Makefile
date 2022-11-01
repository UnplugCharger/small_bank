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

peepdb:
	docker exec -it small_bank psql -U root -d small_bank

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/UnplugCharger/small_bank/db/sqlc Store


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test format peepdb server mock