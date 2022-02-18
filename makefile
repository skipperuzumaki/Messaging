postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root messaging_app

dropdb:
	docker exec -it postgres12 dropdb messaging_app

migrateup:
	migrate -path db/migrations -database "postgresql://root@localhost:5432/messaging_app?sslmode=disable;password=secret" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root@localhost:5432/messaging_app?sslmode=disable;password=secret" -verbose down

.PHONY: postgres createdb dropdb migrateup