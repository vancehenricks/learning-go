.PHONY: up down seed

all: up seed

up:
	docker compose up -d
	sleep 2
down:
	docker compose down -v
seed: 
	cat seed.sql | docker exec -i db-1 psql -h localhost -U postgres -f-