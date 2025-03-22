package main

import (
	"github.com/feynmaz/kafkag/config"
	"github.com/feynmaz/kafkag/logger"
	"github.com/feynmaz/kafkag/sarama"
)

func main() {
	cfg, err := config.GetDefault()
	if err != nil {
		panic(err)
	}

	logger := logger.New()
	logger.Info().Msgf("%+v\n", cfg)

	producer, err := sarama.NewProducer(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create producer")
	}

	err = producer.SendMessages()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to send messages")
	}
	logger.Info().Msg("finished sending messages")

	_ = producer.Close()
	logger.Info().Msg("finished closing producer")
}
