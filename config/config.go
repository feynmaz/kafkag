package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppID            string   `json:"app_id" env:"APP_ID"`
	BootstrapServers []string `json:"bootstrap_servers" env:"BOOTSTRAP_SERVERS"`
	TopicName        string   `json:"topic_name" env:"TOPIC_NAME"`
	NumEvents        int      `json:"num_events" env:"NUM_EVENTS"`
}

func GetDefault() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return &cfg, fmt.Errorf("failed to parse: %w", err)
	}

	return &cfg, nil
}
