build:
	docker compose build
up-minio:
	docker-compose up -d minio

up-app:
	docker-compose up -d app

up: up-minio up-app

down:
	docker-compose down

restart: down up

minio-console:
	docker-compose exec minio sh

app-shell:
	docker-compose exec app sh
