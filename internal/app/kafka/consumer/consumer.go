package consumer

import (
	"context"
	"encoding/json"
	"mnc/internal/app/kafka"
	"mnc/internal/app/transfer"
	"mnc/internal/repositories"

	"github.com/rs/zerolog/log"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	ready    chan bool
	repo     repositories.Querier
	transfer transfer.ITransferService
}

func NewKafkaConsumer(repo repositories.Querier, transfer transfer.ITransferService) *KafkaConsumer {
	return &KafkaConsumer{
		ready:    make(chan bool),
		repo:     repo,
		transfer: transfer,
	}
}

// Implement sarama.ConsumerGroupHandler interface
func (c *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *KafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *KafkaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			var transferMsg kafka.TransferMessage
			if err := json.Unmarshal(message.Value, &transferMsg); err != nil {
				log.Error().Err(err).Send()
				continue
			}

			log.Info().Msgf("Received transfer message: %v", transferMsg)

			err := c.transfer.ReceivedTransfer(context.Background(), transferMsg)
			if err != nil {
				log.Error().Err(err).Send()
				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
