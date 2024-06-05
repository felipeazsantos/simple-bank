createdb:
	docker exec -it postgres12_sb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12_sb dropdb simple_bank

postgres:
	docker run --name postgres12_sb -p 5442:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v pgdata:/var/lib/postgresql/data -d postgres:12-alpine

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test