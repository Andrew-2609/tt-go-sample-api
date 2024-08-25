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

# Testing

test-unit:
	go clean -testcache && ENV=test go test -short ./...

test-unit-v:
	go clean -testcache && ENV=test go test -v -short ./...

test-integration:
	go clean -testcache && ENV=test go test -p 1 ./... -run Integration

test-integration-v:
	go clean -testcache && ENV=test go test -v -p 1 ./... -run Integration

.PHONY: test-unit test-unit-v

# Mocking

update-mocks:
	@mockgen -source=./domain/entity/employee_contract.go -destination=./mock/employee_contract_mock.go -package=mock

.PHONY: update-mocks