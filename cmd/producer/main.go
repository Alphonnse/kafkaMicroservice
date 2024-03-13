package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderPlacer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewOrderProducer(p *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlacer) placeOrder(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload, // Best way to do this is a protobuf
	},
		op.deliverych,
	)
	if err != nil {
		log.Fatal(err)
	}

	<- op.deliverych

	fmt.Printf("Placed order on the queue: %s\n", format)

	return nil
}

func main() {
	// Producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "something",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	op := NewOrderProducer(p, "HVSE")

	for i:=0;i<1000;i++{
		if err := op.placeOrder("maret", i+1); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second*3)
	}
}
