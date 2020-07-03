package test

import (
	pb "emitter/proto/destroyer_pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type TestServer struct{}

//InitializeMockServer Initialized a constructor  Of TestServer struct
func InitializeMockServer() *TestServer {
	return &TestServer{}
}

func MockServer(server pb.CommunicationServer) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterCommunicationServer(s, server)

	log.Println("Starting Server On Port:")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
