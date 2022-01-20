gen:
		rm -rf rpc
		mkdir rpc
		protoc ./proto/* --go_out=. --go-grpc_out=. --twirp_out=.

# for more info about the protoc CLI, see docs: https://grpc.io/docs/languages/go/quickstart/

serve:
		go run main.go

mongo:
		export DB_ENGINE="mongo" && \
    go run main.go

postgres:
		export DB_ENGINE="postgres" && \
		go run main.go

test:
		go test ./server/server_test.go