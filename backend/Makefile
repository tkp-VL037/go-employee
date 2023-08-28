gql:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.36 && go get github.com/99designs/gqlgen/codegen@v0.17.36 && go get github.com/99designs/gqlgen@v0.17.36 && go run github.com/99designs/gqlgen generate

protocc:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/*.proto

run:
	go run server.go
