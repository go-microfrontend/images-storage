up-minio:
	sudo docker-compose up -d minio

up-app:
	sudo docker-compose up -d app

up: up-minio up-app

down:
	sudo docker-compose down

restart: down up

minio-console:
	sudo docker-compose exec minio sh

app-shell:
	sudo docker-compose exec app sh
