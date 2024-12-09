.PHONY: build up down

build:
	docker build . --file Dockerfile -t app:latest
    docker-compose build

up:
	docker build . --file Dockerfile -t app:latest
    docker-compose up

down:
    docker-compose down
