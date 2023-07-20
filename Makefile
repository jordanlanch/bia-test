DB_NAME=bia-test-db
DB_PORT=5432
MIGRATION_DIR=storage/migrations
ROUTE="host=localhost user=postgres password=password dbname=${DB_NAME} port=${DB_PORT} sslmode=disable"

define setup_env
    $(eval ENV_FILE := .env.test_integration)
    @echo " - setup env $(ENV_FILE)"
    $(eval include .env.test_integration)
    $(eval export)
endef

## get extra arguments and filter out commands from args
args = $(filter-out $@,$(MAKECMDGOALS))

.PHONY: all test unit_test e2e_test

test:

	echo "Starting test environment"
	$(call setup_env)
	make unit_test
	make e2e_test

unit_test:
	echo "/////////////////////////////////Starting unit test environment/////////////////////////////////"
	cd ./restserver && go test ./... -coverprofile coverage.out -covermode count &&  go tool cover -func coverage.out | grep total | awk '{print $3}'
	cd ./services/anti_fraud/manager && go test ./... -coverprofile coverage.out -covermode count &&  go tool cover -func coverage.out | grep total | awk '{print $3}'
	cd ./services/hard_internal_rules/manager && go test ./... -coverprofile coverage.out -covermode count && go tool cover -func coverage.out | grep total | awk '{print $3}'
	cd ./restserver && go tool cover -html=coverage.out -o cover.html
	cd ./services/anti_fraud/manager && go tool cover -html=coverage.out -o cover.html
	cd ./services/hard_internal_rules/manager && go tool cover -html=coverage.out -o cover.html

goose_install:
	go install github.com/pressly/goose/v3/cmd/goose@v3.5.3

e2e_test: goose_install
	echo "/////////////////////////////////Deleting fixtures/////////////////////////////////"
	rm -rf ./test/fixtures
	echo "/////////////////////////////////Starting E2E Test/////////////////////////////////"
	docker compose -f docker-compose-test.yaml up --build -d
	sh -c 'sleep 5 &&  goose -dir ${MIGRATION_DIR} postgres "host=localhost user=postgres dbname=deuna-guard-test port=5470 sslmode=disable" up &&  goose -dir seeds postgres "host=localhost user=postgres dbname=deuna-guard-test port=5470 sslmode=disable" up'
	cd ./test &&  go test ./...
	docker compose -f docker-compose-test.yaml down
	echo "/////////////////////////////////Ending E2E Test/////////////////////////////////"

## default that allows accepting extra args
%:
    @:

.PHONY: migration
migration:
	goose -dir ${MIGRATION_DIR} create $(call args,defaultstring) sql
.PHONY: migration
migration-go:
	goose -dir ${MIGRATION_DIR} create $(call args,defaultstring) go

migrate-status:
	goose -dir ${MIGRATION_DIR} postgres ${ROUTE} status

migrate-up:
	goose -dir ${MIGRATION_DIR} postgres ${ROUTE} up
migrate-seeds:
	./seeds/goose-custom -dir seeds up -dbstring ${ROUTE}

migrate-down:
	goose -dir ${MIGRATION_DIR} postgres ${ROUTE} down

migrate-rollback:
	goose -dir ${MIGRATION_DIR} postgres ${ROUTE} reset

migrate-reset:
	goose -dir ${MIGRATION_DIR} postgres ${ROUTE} reset

mocks:
	mockery --dir=domain --output=domain/mocks --outpkg=mocks --all
