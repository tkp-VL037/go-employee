gql:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.36 && go get github.com/99designs/gqlgen/codegen@v0.17.36 && go get github.com/99designs/gqlgen@v0.17.36 && go run github.com/99designs/gqlgen generate

proto:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. common/model/*.proto

run_db:
	docker run --name postgres-go-employee -p 5432:5432  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e -d postgres:15-alpine

create_db:
	docker exec -it postgres-go-employee createdb -U postgres "goemployee"