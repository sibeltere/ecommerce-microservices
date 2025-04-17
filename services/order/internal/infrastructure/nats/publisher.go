package natsinfra

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	connectionUrl string
}

func NewNatsPublisher(connectionUrl string) *NatsPublisher {
	return &NatsPublisher{connectionUrl: connectionUrl}
}

func (p *NatsPublisher) Publish(subject string, data any) (bool, error) {

	nc, err := nats.Connect(p.connectionUrl)
	if err != nil {
		log.Fatalf("Could not connect to nats %v", err)
		return false, err
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Could not connect to jets %v", err)
		return false, err
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",             // Stream adı
		Subjects: []string{"orders.*"}, // Konuları
		Storage:  nats.FileStorage,     // Diskte kalıcı
	})

	publishedData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Order model is not converted to JSON %v", err)
		return false, err
	}

	ack, err := js.Publish(subject, publishedData)
	if err != nil {
		log.Fatalf("JetStream publishing error: %v", err)
		return false, err
	}

	log.Printf("Message published on JetStream Stream: %s, Seq: %d", ack.Stream, ack.Sequence)

	return true, nil
}
