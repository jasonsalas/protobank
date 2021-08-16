# Protobank - server-side gRPC streaming with Go

## This is an implementation of the super-cool demo of gRPC server-side streaming by [inanzzz](http://www.inanzzz.com/index.php/post/w027/creating-a-server-side-grpc-streaming-with-golang)

This project uses protocol buffers and Go to create a simple virtual bank that streams a customer's account transactions to the client in the terminal.

## Usage patterns
### 1. From command line
- Start the server: `go run --race cmd/server`
- Start the client: `go run --race cmd/client`

### 2. From Docker locally
- Build the image in Docker: `docker build -t protobank .`
- Spin-up a container: `docker run --name protobank -p 50051:50051 -d protobank`
- Start the client `go run --race ./cmd/client`

### 3. From Docker Hub
- Pull the [cloud container image](https://hub.docker.com/repository/docker/jasonsalas/protobank): `docker pull jasonsalas/protobank:v1.0`
- Spin-up a container: `docker run --name protobank -p 50051:50051 -d jasonsalas/protobank:v1.0`
- Compile the client: `go build ./cmd/client`, then `./cmd/client`

## Result:
On the server

![On the server](https://github.com/jasonsalas/protobank/blob/main/server.JPG?raw=true)

On the client

![On the client](https://github.com/jasonsalas/protobank/blob/main/client.JPG?raw=true)