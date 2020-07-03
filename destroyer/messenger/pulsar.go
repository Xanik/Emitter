package messenger

import (
	"context"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Connect instantiate Pulsar a client
func Connect(url string) (pulsar.Client, error) {
	log.Printf("\nCreating new Pulsar connection")

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

// Producer publishes messages to pulsar
func Producer(client pulsar.Client, topic string, payload []byte) error {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: payload,
	})

	defer producer.Close()

	if err != nil {
		log.Println("Failed to publish message", err)
		return err
	}
	log.Println("Published message")
	return nil
}
