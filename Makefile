tidy:
	docker compose run --rm web go mod tidy

# For docker-compose workflow
up:
	docker compose up

stop:
	docker compose stop
