include .env
export 
#Challenge Makefile
start:
	go run main.go

check:
#TODO: include command to test the code and show the results

#setup:
postgres:
	docker run --name postgres971224 -p $(PORT):$(PORT) -e POSTGRES_USER=$(USER) -e POSTGRES_PASSWORD=$(PASSWORD) -d postgres:12-alpine

createdb:
	docker exec -it postgres971224 createdb --username=$(USER) --owner=$(USER) $(DBNAME)

dropdb:
	docker exec -it postgres971224 dropdb $(DBNAME)

migrateup:
	migrate -path db/migration -database $(MY_MIGRATE_DATABASE) -verbose up

migratedown:
	migrate -path db/migration -database $(MY_MIGRATE_DATABASE) -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown 