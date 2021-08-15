.PHONY: compile

compile:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative pkg/protobuf/bank/*.proto