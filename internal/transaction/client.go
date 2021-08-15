package transaction

import (
	"context"
	"io"
	"log"

	bank "github.com/jasonsalas/protobank/pkg/protobuf/bank"
	"google.golang.org/grpc"
)

type Client struct {
	client bank.TransactionServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: bank.NewTransactionServiceClient(conn),
	}
}

func (c Client) Fetch(ctx context.Context, accountID string) error {
	stream, err := c.client.Fetch(ctx, &bank.FetchRequest{AccountId: accountID})
	if err != nil {
		return err
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("> ALL DONE!")
				return nil
			}

			return err
		}

		trx := response.GetTransaction()
		log.Println("TIME:", trx.GetTime().String())
		log.Println("OPERATION:", trx.GetOperation().String())
		log.Println("AMOUNT:", trx.GetAmount())
		log.Println("------")
	}
}
