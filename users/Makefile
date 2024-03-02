build:
	go build github.com/Prokopevs/kanban/users/cmd/server
run:
	PG_CONN=postgres://postgres:2409@localhost:5432/laniakeadb?sslmode=disable \
	GRPC_SERVER_ADDR=":8089" \
	HTTP_SERVER_ADDR=":3000" \
	go run cmd/server/env.go cmd/server/main.go
	
