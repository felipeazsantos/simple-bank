postgres:
	docker run --name postgres_simple_bank -p 5442:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v /home/felipe/Documentos/Dev-pessoal/golang/simple-bank/data:/var/lib/postgresql/data -d postgres:12-alpine
createdb:
	docker exec -it postgres_simple_bank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres_simple_bank dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

fixdata:
	sudo chmod 777 -R data

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test