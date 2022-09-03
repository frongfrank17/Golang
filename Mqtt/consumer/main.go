package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("SERVICE CONSUMER")
	fmt.Println("SERVICE PRODUCER")
	server := []string{"localhost:9092"}
	topic := "my-queue-1"
	consumer, err := sarama.NewConsumer(server, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("RUNNING Consume")
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	fmt.Println("Consumer start.")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err) 
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}

}
