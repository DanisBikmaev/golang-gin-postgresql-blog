.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root blog
dropdb:
	docker exec -it postgres15 dropdb blog
migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/blog?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/blog?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...