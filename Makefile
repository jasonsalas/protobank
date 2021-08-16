.PHONY: compile

compile:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative pkg/protobuf/bank/*.proto

binary:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o protobank-server ./cmd/server/main.go