# Protobank - server-side gRPC streaming with Go

## This is an implementation of the super-cool demo of gRPC server-side streaming by [inanzzz](http://www.inanzzz.com/index.php/post/w027/creating-a-server-side-grpc-streaming-with-golang)

This project uses protocol buffers and Go to create a simple virtual bank that streams a customer's account transactions to the client in the terminal.

### Usage
- Start the server: `go run --race cmd/server`
- Start the client: `go run --race cmd/client`
