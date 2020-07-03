gen-proto:
	cd proto && protoc --go_out=plugins=grpc:./deathstar_pb deathstar.proto

test:
	cd db/__test__ && go test ./...