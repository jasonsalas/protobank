package transaction

import (
	"time"

	bank "github.com/jasonsalas/protobank/pkg/protobuf/bank"
)

type transaction struct {
	time      time.Time
	operation bank.Transaction_Operation
	amount    float64
}

var transactions = map[string][]transaction{
	"ACC-1": {
		{
			time:      time.Now().Add(time.Second),
			operation: bank.Transaction_Credit,
			amount:    10.00,
		},
		{
			time:      time.Now().Add(time.Millisecond),
			operation: bank.Transaction_Credit,
			amount:    20.99,
		},
		{
			time:      time.Now().Add(time.Hour),
			operation: bank.Transaction_Debit,
			amount:    30.62,
		},
		{
			time:      time.Now().Add(time.Minute),
			operation: bank.Transaction_Credit,
			amount:    40,
		},
		{
			time:      time.Now().Add(time.Minute),
			operation: bank.Transaction_Debit,
			amount:    50.55,
		},
		{
			time:      time.Now().Add(time.Minute),
			operation: bank.Transaction_Debit,
			amount:    60.60,
		},
	},
}
