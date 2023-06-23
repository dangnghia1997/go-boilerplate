tidy:
	docker compose run --rm web go mod tidy

# For docker-compose workflow
up:
	docker compose up

stop:
	docker compose stop

# Number of migrations - this is optionally used on up and down commands
N?=

MYSQL_USER ?= root
MYSQL_PASSWORD ?= P@ssw0rd
MYSQL_HOST ?= 127.0.0.1
MYSQL_DATABASE ?= go_crud
MYSQL_PORT ?= 3306
MYSQL_DSN ?= $(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)

migrate-create:
	@ migrate create -ext sql -dir migrations/ -seq $(NAME)

migrate-up:
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations up $(N)

migrate-down:
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations down $(N)

drop-db:
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations drop

migration-version:
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path migrations version
