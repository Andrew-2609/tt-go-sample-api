# Running

run:
	go run main.go

.PHONY: run

# Database

sqlc:
	sqlc generate
	
create-migration:
	migrate create -ext sql -dir external/rdb/migration -seq $(name)

migrate-up:
	migrate -path external/rdb/migration -database "postgresql://ndrew:db_pass@localhost:5432/db_name?sslmode=disable" -verbose up $(n)

migrate-down:
	migrate -path external/rdb/migration -database "postgresql://ndrew:db_pass@localhost:5432/db_name?sslmode=disable" -verbose down $(n)

.PHONY: sqlc create-migration migrate-up migrate-down