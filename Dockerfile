FROM golang:1.18 AS build

WORKDIR /go/src/github.com/MitoVeli/math_grpc_server
COPY . .

RUN go build -o /go/bin/math_grpc_server

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/math_grpc_server /go/bin/math_grpc_server

ENV APP_PORT=8080
ENV GRPC_PORT=50052

ENTRYPOINT ["/go/bin/math_grpc_server"]

