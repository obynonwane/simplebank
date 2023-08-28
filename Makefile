postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres 

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:gobank@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:gobank@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	cd db/sqlc && go test -v

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
