package main

import (
	"emitter/client_test/proto/pb"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	NewClient()
	time.Sleep(2000 * time.Minute)

}

// NOTE RUN SERVICE WITH DOCKER BEFORE RUNNING THIS TEST
// New Storage Client Connection Created
func NewClient() pb.CommunicationClient {
	port := ":8882"

	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to Serve: %v", err)
	}

	client := pb.NewCommunicationClient(conn)

	fmt.Println("connected to destroyer service")

	return client

}
