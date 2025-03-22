package sarama

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/feynmaz/kafkag/config"
	"github.com/feynmaz/kafkag/logger"
)

type Producer struct {
	cfg    *config.Config
	logger *logger.Logger

	producer sarama.SyncProducer
}

func NewProducer(cfg *config.Config, logger *logger.Logger) (*Producer, error) {
	logger.Info().Msg("Creating Kafka Producer...")

	// Step 1. Set producer confuguration
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Version = sarama.V3_9_0_0
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.ClientID = cfg.AppID

	// Step 2. Create producer instance
	producer, err := sarama.NewSyncProducer(cfg.BootstrapServers, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka Producer: %w", err)
	}

	return &Producer{
		cfg:    cfg,
		logger: logger,

		producer: producer,
	}, nil
}

func (p Producer) SendMessages() error {
	// Step 3. Send all messages
	// NOTE: 1MB ~ 6k messages
	for i := range p.cfg.NumEvents {
		msg := &sarama.ProducerMessage{
			Topic: p.cfg.TopicName,
			Key:   sarama.StringEncoder(fmt.Sprintf("%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Simple message - %d", i)),
		}
		_, _, err := p.producer.SendMessage(msg)
		if err != nil {
			return fmt.Errorf("failed to produce message: %v", err)
		}
	}
	return nil
}

func (p Producer) Close() error {
	// Step 4. Close the producer
	return p.producer.Close()
}
