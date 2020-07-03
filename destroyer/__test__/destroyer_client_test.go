package test

import (
	pb "emitter/proto/destroyer_pb"

	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
)

// Initialize Server Service
func init() {
	go MockServer(InitializeMockServer())
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

// New Storage Client Connection Created
func newClient() pb.CommunicationClient {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to Serve: %v", err)
	}
	client := pb.NewCommunicationClient(conn)
	return client
}

func TestAcquireTargets(t *testing.T) {
	client := newClient()

	mockData := mockTarget{
		ID:        "01EBP4DP4VECW8PHDJJFNEDVKE",
		Message:   "send a message",
		CreatedOn: "2020-06-25T16:23:37.720Z",
	}

	req := &pb.EventMessage{
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
	}
	res, err := client.AcquireTargets(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if res.GetId() != req.GetId() {
		t.Errorf("%v is not equal to %v", res, req)
	}
	if res.GetName() != req.GetName() {
		t.Errorf("%v is not equal to %v", res, req)
	}
	if res.GetCreatedOn() != req.GetCreatedOn() {
		t.Errorf("%v is not equal to %v", res, req)
	}
}

func TestListTargets(t *testing.T) {
	client := newClient()

	res, err := client.ListTargets(context.Background(), &pb.TargetRequestMessage{
		Id: "01EBP4DP4VECW8PHDJJFNEDVKE",
	})

	if err != nil {
		t.Error(err)
	}
	if len(res.GetTargetResponses()) == 0 {
		t.Errorf("Invalid Result: %v", res)
	}
}
