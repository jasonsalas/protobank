package transaction

import (
	"log"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes"
	bank "github.com/jasonsalas/protobank/pkg/protobuf/bank"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var sleeper = []int{0, 1, 2, 3, 4}

type Server struct {
	transactions map[string][]transaction
}

func NewServer() Server {
	return Server{
		transactions: transactions,
	}
}

func (s Server) Fetch(request *bank.FetchRequest, stream bank.TransactionService_FetchServer) error {
	rand.Seed(time.Now().UnixNano())
	log.Println("fetching transactions for account:", request.GetAccountId())

	trxs := s.transactions[request.GetAccountId()]

	for _, trx := range trxs {
		ts, err := ptypes.TimestampProto(trx.time)
		if err != nil {
			return status.Errorf(codes.Internal, "fetch: invalid time: %v", err)
		}

		if err := stream.Send(&bank.FetchResponse{
			Transaction: &bank.Transaction{
				Time:      ts,
				Operation: trx.operation,
				Amount:    trx.amount,
			},
		}); err != nil {
			return status.Errorf(codes.Internal, "fetch: unexpected stream: %v", err)
		}
		sleep := rand.Intn(len(sleeper))
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	log.Println("completed")
	return nil
}
