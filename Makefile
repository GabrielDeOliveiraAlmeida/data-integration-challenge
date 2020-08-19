#Challenge Makefile
start:
	go run main.go

check:
#TODO: include command to test the code and show the results

#setup:
postgres:
	docker run --name postgres971224 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres971224 createdb --username=root --owner=root companiesDb

dropdb:
	docker exec -it postgres971224 dropdb companiesDb

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@192.168.99.100:5432/companiesDb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@192.168.99.100:5432/companiesDb?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown 