package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// There also can be different froup IDs to read same mesages not once
func main() {
	// Processor
	topic := "HVSE"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest",
	})
	defer consumer.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		consumer.Close()
		os.Exit(1)
	}()

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Data team reading order: %s\n", string(e.Value))
		case *kafka.Error:
			fmt.Printf("%v\n", e)
		}
	}
}
