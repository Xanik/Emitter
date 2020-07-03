package test

import (
	"context"
	"emitter/destroyer/proto/pb"
)

type mockTarget struct {
	ID        string
	Message   string
	CreatedOn string
}

func (s *TestServer) AcquireTargets(ctx context.Context, m *pb.EventMessage) (*pb.EventMessage, error) {
	return &pb.EventMessage{
		Id:        m.GetId(),
		Name:      m.GetName(),
		Data:      m.Data,
		CreatedOn: m.GetCreatedOn(),
	}, nil
}

func (s *TestServer) ListTargets(ctx context.Context, m *pb.TargetRequestMessage) (*pb.TargetResponseMessage, error) {
	mockData := mockTarget{
		ID:        m.GetId(),
		Message:   "send a message",
		CreatedOn: "2020-06-25T16:23:37.720Z",
	}
	return &pb.TargetResponseMessage{
		TargetResponses: []*pb.TargetResponse{
			{
				Id:        m.GetId(),
				Message:   mockData.Message,
				CreatedOn: mockData.CreatedOn,
			},
		},
	}, nil
}
