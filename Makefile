postgres:
	docker run --name postgres_simple_bank -p 5442:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v /home/felipe/Documentos/Dev-pessoal/golang/simple-bank/data:/var/lib/postgresql/data -d postgres:12-alpine
createdb:
	docker exec -it postgres_simple_bank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres_simple_bank dropdb simple_bank

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

fixdata:
	sudo chmod 777 -R data

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/felipeazsantos/simple_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test fixdata server mock