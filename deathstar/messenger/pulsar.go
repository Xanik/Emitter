package messenger

import (
	"context"
	"emitter/deathstar/proto/deathstar_pb"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Connect instantiate Pulsar a client
func Connect() (pulsar.Client, error) {
	log.Printf("\nCreating new Pulsar connection")

	url := fmt.Sprintf("pulsar://%v:6650", os.Getenv("PULSAR"))

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
		return nil, err
	}

	return client, nil
}

// Consumer consumes messages from pulsar and store in db
func Consumer(client pulsar.Client, topic string) (deathstar_pb.EventMessage, error) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	var event deathstar_pb.EventMessage

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
		return deathstar_pb.EventMessage{}, err
	}

	fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
		msg.ID(), string(msg.Payload()))

	consumer.Ack(msg)

	err = json.Unmarshal(msg.Payload(), &event)
	if err != nil {
		return deathstar_pb.EventMessage{}, err
	}

	return event, nil
}
