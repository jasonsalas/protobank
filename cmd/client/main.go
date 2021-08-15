package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jasonsalas/protobank/internal/transaction"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting client")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	txClient := transaction.NewClient(conn)
	if err := txClient.Fetch(context.Background(), "ACC-1"); err != nil {
		log.Fatalln(err)
	}
}
