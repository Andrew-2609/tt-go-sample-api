# Running

run:
	docker compose up -d --remove-orphans && docker logs -f tt-go-sample-api

run-build:
	docker compose up --build -d --remove-orphans && docker logs -f tt-go-sample-api

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

up-test:
	docker compose --env-file .env.test -f docker-compose.testing.yml up -d --build --remove-orphans

test-unit:
	go clean -testcache && ENV=test go test -short ./...

test-unit-v:
	go clean -testcache && ENV=test go test -v -short ./...

# ⚠️ Integrations Tests will likely only work in a dockerized environment ⚠️

test-integration:
	go clean -testcache && ENV=test go test -p 1 ./... -run Integration

test-integration-v:
	go clean -testcache && ENV=test go test -v -p 1 ./... -run Integration

test: test-unit test-integration

test-v: test-unit-v test-integration-v

coverage:
	@docker compose --env-file .env.test -f docker-compose.testing.yml up -d --build --remove-orphans
	@docker exec test-runner ./test_coverage_runner.sh
	@docker compose -f docker-compose.testing.yml down
	
coverage-html: coverage
	go tool cover -html=tmp/coverage.out -o tmp/coverage.html
	xdg-open tmp/coverage.html

.PHONY: up-test test-unit test-unit-v test-integration test-integration-v test test-v coverage coverage-html

# Mocking

update-mocks:
	@mockgen -source=./domain/entity/employee_contract.go -destination=./mock/employee_contract_mock.go -package=mock

.PHONY: update-mocks