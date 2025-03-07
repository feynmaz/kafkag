package confluent

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/feynmaz/kafkag/config"
	"github.com/feynmaz/kafkag/logger"
)

// Sending messages is a 4-step process

type Producer struct {
	cfg    *config.Config
	logger *logger.Logger

	producer *kafka.Producer
}

func NewProducer(cfg *config.Config, logger *logger.Logger) (*Producer, error) {
	logger.Info().Msg("Creating Kafka Producer...")

	// Step 1. Set producer confuguration
	kafkaCfg := &kafka.ConfigMap{
		"client.id": cfg.AppID,
		// It is recommended to provide 2-3 brokers in multi-brokers cluster
		// This will help in case first broker in the list is down
		"bootstrap.servers": cfg.BootstrapServers,
	}

	// Step 2. Create producer instance
	producer, err := kafka.NewProducer(kafkaCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create  Kafka Producer")
	}

	return &Producer{
		cfg:      cfg,
		logger:   logger,
		producer: producer,
	}, nil
}

func (p Producer) SendMessages() error {
	// Step 3. Send all messages
	for i := range p.cfg.NumEvents {
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic: &p.cfg.TopicName,
			},
			Key:   []byte(fmt.Sprintf("%d", i)),
			Value: []byte(fmt.Sprintf("Simple message - %d", i)),
		}
		err := p.producer.Produce(msg, nil)
		if err != nil {
			return fmt.Errorf("failed to produce message: %v", err)
		}
	}
	return nil
}

func (p Producer) Close(timeout time.Duration) {
	// Step 4. Close the producer
	p.producer.Flush(int(timeout.Milliseconds()))
	p.producer.Close()
}
