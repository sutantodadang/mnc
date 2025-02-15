package infrastructure

import (
	"os"

	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Producer sarama.SyncProducer
	Consumer sarama.ConsumerGroup
}

func InitKafka() (*KafkaConfig, error) {

	// producer config
	producerCfg := sarama.NewConfig()
	producerCfg.Producer.Return.Successes = true
	producerCfg.Producer.RequiredAcks = sarama.WaitForAll
	producerCfg.Producer.Retry.Max = 10

	// consumer config
	consumerCfg := sarama.NewConfig()
	consumerCfg.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	consumerCfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_ADDR")}, producerCfg)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumerGroup([]string{os.Getenv("KAFKA_ADDR")}, os.Getenv("KAFKA_GROUP_ID"), consumerCfg)
	if err != nil {
		return nil, err
	}

	return &KafkaConfig{
		Producer: producer,
		Consumer: consumer,
	}, nil

}
