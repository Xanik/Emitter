gen-proto:
	cd proto && protoc --go_out=plugins=grpc:./destroyer_pb destroyer.proto

test:
	cd __test__ && go test ./...
	