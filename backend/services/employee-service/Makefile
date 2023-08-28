include .env

run_db:
	docker run --name postgres-go-employee -p 5432:5432  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:15-alpine

create_db:
	docker exec -it postgres-go-employee createdb -U postgres "goemployee"

migrate_up:
	migrate -path "db/migrations" -database "$(POSTGRESQL_URL)" up

migrate_down:
	migrate -path "db/migrations" -database "$(POSTGRESQL_URL)" down

run_redis:
	docker run --name redis-go-employee -p 6379:6379 -d redis:7.2.0-alpine