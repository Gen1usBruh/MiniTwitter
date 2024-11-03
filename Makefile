migration_path := internal/storage/postgres/migration
home_path := /home/poset
migrateinit:
	$(home_path)/bin/migrate create -ext sql -dir $(migration_path) -seq init_schema
migrateup:
	$(home_path)/bin/migrate -path $(migration_path) -database "postgresql://amirshox:aboba2298@localhost:5433/twitterdb?sslmode=disable" -verbose up
migratedown:
	$(home_path)/bin/migrate -path $(migration_path) -database "postgresql://amirshox:aboba2298@localhost:5433/twitterdb?sslmode=disable" -verbose down
connectdb:
	docker exec -it postgres15.8 psql -U amirshox -d twitterdb
createdb:
	docker exec -it postgres15.8 createdb --username=amirshox --owner=amirshox twitterdb
dropdb:
	docker exec -it postgres15.8 dropdb -U amirshox twitterdb
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: connectdb createdb dropdb migrateup migratedown migrateinit sqlc test
