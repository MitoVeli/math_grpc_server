# math_grpc_server
Math operations grpc server; gets operation sign and variables from the [math_grpc_client](https://github.com/MitoVeli/math_grpc_client), and returns the result of the operation.

## Run
It is not needed to run the application as the services do not have any external connections. They can be reached offline directly via Grpc.

However, the application is already dockerized just for demonstration purposes and it can be run by the command below;

You can run the project by using
-   `docker-compose up --build`

## Test
Tests can be run by the following command;
-   `go test ./... -v`

### Comments
Http server in main.go is also added just for demonstration purposes. In case the application is run by docker, http server keeps the application up and running.