package main

import (
	"context"
	"emitter/deathstar/db"
	"emitter/deathstar/messenger"
	"emitter/deathstar/proto/deathstar_pb"

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

	conf := viper.GetStringMapString("deathstar")
	// Connect to DB
	DB, err := db.Connect()
	if err != nil {
		fmt.Errorf("Can't connect to db: %v", err)
	}
	// Connect To Pulsar
	pulse, err := messenger.Connect()
	if err != nil {
		fmt.Errorf("Can't connect to pulsar: %v", err)
	}
	// Run GRPC
	go newService(InitializeServer(pulse, DB), ":"+conf["grpc_port"])

	for {
		// Get From Pulsar Here And Store in DB
		event, err := messenger.Consumer(pulse, conf["topic"])
		if err != nil {
			fmt.Errorf("Pulsar Error: %v", err)
			return
		}
		_, err = DB.SaveTarget(&event)
		if err != nil {
			fmt.Errorf("Db Error: %v", err)
			return
		}
	}
}

//Server Struct
type Server struct {
	mutex *sync.RWMutex
	db    db.DB
	// Pulsar Goes Here......
	pulsar pulsar.Client
}

//InitializeServer Initialized a constructor  Of Server struct
func InitializeServer(pulsar pulsar.Client, DB *db.DB) *Server {
	return &Server{mutex: &sync.RWMutex{}, db: *DB, pulsar: pulsar}
}

// StoreTarget takes in Event and Stores it's Targets
func (s *Server) StoreTarget(ctx context.Context, m *deathstar_pb.EventRequestMessage) (*deathstar_pb.EventResponseMessage, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata Received: %v\n", md)
	}

	var event *deathstar_pb.EventMessage

	s.mutex.RLock()
	_, err := s.db.SaveTarget(event)
	if err != nil {
		return nil, err
	}
	s.mutex.RUnlock()

	return &deathstar_pb.EventResponseMessage{
		Id:        event.Id,
		Message:   event.Name,
		CreatedOn: event.CreatedOn,
	}, nil
}

func newService(server deathstar_pb.CommunicationServer, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	deathstar_pb.RegisterCommunicationServer(s, server)

	log.Println("Starting Server On Port:" + port)

	e := s.Serve(lis)
	if e != nil {
		panic(err)
	}
}
