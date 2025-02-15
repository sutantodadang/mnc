package kafka

import (
	"encoding/json"
	"os"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(producer sarama.SyncProducer) *KafkaProducer {
	return &KafkaProducer{
		producer: producer,
	}
}

func (k *KafkaProducer) PublishTransfer(message TransferMessage) (err error) {

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC"),
		Value: sarama.StringEncoder(jsonMessage),
		Key:   sarama.StringEncoder(message.TransferID),
	})
	if err != nil {
		return err
	}

	return
}
