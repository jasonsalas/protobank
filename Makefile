.PHONY: compile

compile:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative pkg/protobuf/bank/transaction.proto

binary:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o protobank-server ./cmd/server/main.go

docker:
	go build -o ./cmd/client/protobank-client ./cmd/client/main.go
	docker run --name protobank -p 50051:50051 -d jasonsalas/protobank:v1.0
	./cmd/client/protobank-client