gen-proto:
	cd proto && protoc --go_out=plugins=grpc:./pb destroyer.proto

test:
	cd __test__ && go test ./...
	cd db/__test__ && go test ./...
	