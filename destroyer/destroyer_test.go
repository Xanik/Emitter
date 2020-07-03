package main

import (
	"context"
	"emitter/destroyer/proto/pb"
	"log"
	"testing"

	"google.golang.org/grpc"
)

// NOTE RUN SERVICE WITH DOCKER BEFORE RUNNING THIS TEST
// New Storage Client Connection Created
func newClient() pb.CommunicationClient {
	port := ":8882"

	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to Serve: %v", err)
	}

	client := pb.NewCommunicationClient(conn)
	return client

}

// TestAcquireTargets Function
func TestAcquireTargets(t *testing.T) {
	type mockTarget struct {
		ID        string
		Message   string
		CreatedOn string
	}

	mockData := mockTarget{
		ID:        "01EBP4DP4VECW8PHDJJFNEDVKE",
		Message:   "send a message",
		CreatedOn: "2020-06-25T16:23:37.720Z",
	}

	res, err := newClient().AcquireTargets(context.Background(), &pb.EventMessage{
		Id:   "01EBP4DP4VECW8PHDJJFNEDVKE",
		Name: "targets.acquired",
		Data: []*pb.TargetResponse{
			{
				Id:        "01EBP4DP4VECW8PHDJJFNEDVKE",
				Message:   mockData.Message,
				CreatedOn: mockData.CreatedOn,
			},
		},
		CreatedOn: "2020-06-25T16:23:37.720Z",
	})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}

// TestListTargets Function
func TestListTargets(t *testing.T) {
	res, err := newClient().ListTargets(context.Background(), &pb.TargetRequestMessage{
		Id: "01EBP4DP4VECW8PHDJJFNEDVKE",
	})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}
