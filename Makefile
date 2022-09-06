.PHONY: clean build run_api run_import migrations_up migrations_down generate_models test

psql = "postgres://postgres:pwd@localhost:5432/postgres?sslmode=disable"

clean:
	-rm run

build: clean
	golangci-lint run
	go generate
	go build -race -o run

run_api: build
	./run api

run_import: build
	./run import

migrations_up:
	goose -dir migrations postgres $(psql) up

migrations_down:
	goose -dir migrations postgres $(psql) down

generate_models:
	sqlboiler -c=sqlboiler.toml --wipe --no-auto-timestamps psql

test:
	go test ./... -coverprofile .testCoverage.out

test_unit:
	go test $(go list ./... | grep -v /tests/)

coverage:
	go tool cover -html=.testCoverage.out