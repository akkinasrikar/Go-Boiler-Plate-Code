.PHONY: generate-mocks migrateup migratedown

generate-mocks:
	go generate ./...

migrateup:
	migrate -path database/migration -database "postgres://postgres:postgres@127.0.0.1:5432/test?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgres://postgres:postgres@127.0.0.1:5432/test?sslmode=disable" -verbose down

forcemigrate:
	migrate -path database/migration -database "postgres://postgres:postgres@localhost:5432/test?sslmode=disable" -verbose force 1

