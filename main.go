package main

import (
	"fmt"
	"time"

	"github.com/feynmaz/kafkag/config"
	"github.com/feynmaz/kafkag/confluent"
	"github.com/feynmaz/kafkag/logger"
)

func main() {
	cfg := config.New()
	fmt.Printf("%+v\n", cfg)

	logger := logger.New()

	producer, err := confluent.NewProducer(cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create producer")
	}

	err = producer.SendMessages()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to send messages")
	}
	logger.Info().Msg("finished sending messages")

	timeout := 2 * time.Second
	logger.Info().Msg("closing producer in 2 seconds")
	producer.Close(timeout)
	logger.Info().Msg("finished closing producer")
}
