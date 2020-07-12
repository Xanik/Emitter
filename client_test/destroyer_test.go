package main

import (
	"context"
	"emitter/client_test/proto/pb"
	"testing"
)

// NOTE RUN SERVICE WITH DOCKER BEFORE RUNNING THIS TEST
// New Storage Client Connection Created

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

	res, err := NewClient().AcquireTargets(context.Background(), &pb.EventMessage{
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
	res, err := NewClient().ListTargets(context.Background(), &pb.TargetRequestMessage{
		Id: "01EBP4DP4VECW8PHDJJFNEDVKE",
	})
	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}
	t.Log(res)
}
