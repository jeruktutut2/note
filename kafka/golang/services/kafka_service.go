package services

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaService interface {
	SendMessage(ctx context.Context, message string) (response string)
}

type KafkaServiceImplementation struct {
	Producer *kafka.Producer
}

func NewKafkaService(producer *kafka.Producer) KafkaService {
	return &KafkaServiceImplementation{
		Producer: producer,
	}
}

func (service *KafkaServiceImplementation) SendMessage(ctx context.Context, message string) (response string) {
	topic := "text-message"
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(message),
	}
	err := service.Producer.Produce(msg, nil)
	if err != nil {
		response = err.Error()
	}
	response = "success"
	return
}
