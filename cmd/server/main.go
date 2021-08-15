package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jasonsalas/protobank/internal/transaction"
	bank "github.com/jasonsalas/protobank/pkg/protobuf/bank"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting server")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	trxServer := transaction.NewServer()

	bank.RegisterTransactionServiceServer(grpcServer, trxServer)
	log.Fatalln(grpcServer.Serve(listener))
}
