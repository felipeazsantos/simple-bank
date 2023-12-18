postgres:
	docker run --name postgres-simple-bank -p 5442:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres-simple-bank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-simple-bank dropdb simple_bank

.PHONY: createdb