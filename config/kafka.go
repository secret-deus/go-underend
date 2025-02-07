package config

import (
	"github.com/segmentio/kafka-go"
)

const (
	topic   = "asset_management"
	brokers = "localhost:9092"
)

func KafkaWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokers},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func KafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokers},
		GroupID:  "asset_management",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
