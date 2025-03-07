package config

type Config struct {
	AppID            string `json:"app_id"`
	BootstrapServers string `json:"bootstrap_servers"`
	TopicName        string `json:"topic_name"`
	NumEvents        int    `json:"num_events"`
}

func New() *Config {
	return &Config{
		AppID:            "HelloProducer",
		BootstrapServers: "localhost:9092;localhost:9093",
		TopicName:        "hello-producer-topic",
		NumEvents:        1_000_000,
	}
}
