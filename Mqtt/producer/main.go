package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("SERVICE PRODUCER")
	server := []string{"localhost:9092"}
	topic := "my-queue-1"
	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("RUNNING PRODUCER")
	defer producer.Close()
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("hello ssdfsdfsss"),
	}
	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Partition=%v , offset=%v", partition, offset)
}
