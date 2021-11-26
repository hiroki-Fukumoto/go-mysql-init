set-up:
	docker network create go-mysql-init-network
	docker-compose -f ./docker-compose.yml build

build:
	docker-compose -f ./docker-compose.yml build

up:
	docker-compose -f ./docker-compose.yml up

up-d:
	docker-compose -f ./docker-compose.yml up -d

down:
	docker-compose -f ./docker-compose.yml down

exec-api:
	make up-d
	docker-compose -f ./docker-compose.yml exec api bash || true
