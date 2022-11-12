# math_grpc_server
Math operations grpc server; gets operation sign and variables from the [math_grpc_client](https://github.com/MitoVeli/math_grpc_client), and returns the result of the operation.

## Run
Application can be run in two ways, from the main application level:

1) Run without docker:

-   `go run ./cmd/math_grpc_server/main.go`

2) Run by docker:

-   `docker-compose up --build`

## Test
Tests can be run by the following command;
-   `go test ./... -v`