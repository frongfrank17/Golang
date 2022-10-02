package service

import "github.com/Shopify/sarama"

type consumerHandler struct {
	eventHandler EventService
}

func NewConsumerHandler(eventHandler EventService) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler}
}

func (obj consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (obj consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		obj.eventHandler.Service(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}

	return nil
}
