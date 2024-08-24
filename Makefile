createdb:
	docker exec -it postgres12_sb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12_sb dropdb simple_bank

postgres:
	docker run --name postgres12_sb -p 5442:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v /home/felipe/Documentos/dev_lab/simple_bank/pgdata:/var/lib/postgresql/data -d postgres:12-alpine

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

fixpgdata:
	sudo chmod 777 -R pgdata

mock:
	mockgen -package mockdb -destination db/mock/Store.go  github.com/felipeazsantos/simple_bank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migrateup1 migratedown migratedown1 sqlc test server mock