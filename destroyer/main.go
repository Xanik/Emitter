package main

import (
	"context"
	"emitter/destroyer/db"
	"emitter/destroyer/messenger"
	"emitter/destroyer/proto/pb"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	//Recover From Panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	conf := viper.GetStringMapString("destroyer")
	// Run GRPC
	newService(InitializeServer(conf["pulsar"], conf["topic"]), ":"+conf["grpc_port"])
}

//Server Struct
type Server struct {
	mutex *sync.RWMutex
	db    db.DB
	// Pulsar Goes Here......
	pulsar pulsar.Client
	topic  string
}

//InitializeServer Initialized a constructor  Of Server struct
func InitializeServer(pulsar string, topic string) *Server {
	// Connect to DB
	DB, err := db.Connect()
	if err != nil {
		log.Printf("Can't connect to db: %v", err)
	}

	// Connect To Pulsar
	pulse, err := messenger.Connect(pulsar)
	if err != nil {
		log.Printf("Can't connect to pulsar: %v", err)
	}

	return &Server{mutex: &sync.RWMutex{}, db: *DB, pulsar: pulse, topic: topic}
}

// AcquireTargets Updates The Status Of The App
func (s *Server) AcquireTargets(ctx context.Context, m *pb.EventMessage) (*pb.EventMessage, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	s.mutex.Lock()
	// Send to Pulsar Here
	body, err := json.Marshal(m)
	if err != nil {
		return &pb.EventMessage{}, err
	}

	messenger.Producer(s.pulsar, s.topic, body)
	s.mutex.Unlock()

	return &pb.EventMessage{
		Id:        m.GetId(),
		Name:      m.GetName(),
		Data:      m.GetData(),
		CreatedOn: m.GetCreatedOn(),
	}, nil
}

// ListTargets takes in Allowed Clients Returns Apps Status
func (s *Server) ListTargets(ctx context.Context, m *pb.TargetRequestMessage) (*pb.TargetResponseMessage, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	s.mutex.RLock()
	// Connect To DB and Return List of Targets Here
	resp, err := s.db.GetAllTargets(m.GetId())
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	s.mutex.RUnlock()

	var targets []*pb.TargetResponse

	for _, v := range resp {
		targets = append(targets, &pb.TargetResponse{
			Id:        v.ID,
			Message:   v.Message,
			CreatedOn: v.CreatedOn,
		})
	}

	return &pb.TargetResponseMessage{
		TargetResponses: targets,
	}, nil
}

func newService(server pb.CommunicationServer, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	pb.RegisterCommunicationServer(s, server)

	log.Println("Starting Server On Port:" + port)

	e := s.Serve(lis)
	if e != nil {
		panic(err)
	}
}
