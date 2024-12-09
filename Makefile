.PHONY: build up down

build:
	docker build . --file Dockerfile -t app:latest
    docker-compose build

up:
	docker build . --file Dockerfile -t app:latest
    docker-compose up
    migrate -path /internal/infrastructure/migrations -database "postgresql://roman:030905romaN@89.46.131.181:5432/shtrav?sslmode=disable" -verbose up


down:
    docker-compose down
