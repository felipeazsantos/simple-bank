createdb:
	docker exec -it postgres12_sb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12_sb dropdb simple_bank

postgres:
	docker run --name postgres12_sb -p 5442:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

migrateup:
	migrate -path db/migration -database “postgresql://root:secret@localhost:5442/simple_bank” --verbose up

migratedown:
	migrate -path db/migration -database “postgresql://root:secret@localhost:5442/simple_bank” --verbose down

.PHONY: createdb, dropdb, postgres, migrateup, migratedown