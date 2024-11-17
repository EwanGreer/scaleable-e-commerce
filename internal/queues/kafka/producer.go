package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer interface {
	Produce([]byte)
}

type KafkaProducer struct {
	client  *kgo.Client
	topic   string
	brokers []string
}

func NewProducer(topic string, brokers []string) *KafkaProducer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ProduceRequestTimeout(5*time.Second),
	)
	if err != nil {
		panic(fmt.Sprintf("Error creating Kafka client: %v", err))
	}

	return &KafkaProducer{
		topic:   topic,
		brokers: brokers,
		client:  client,
	}
}

func (k KafkaProducer) Produce(msg []byte) {
	record := &kgo.Record{
		Topic: k.topic,
		Value: msg,
	}

	k.client.Produce(context.Background(), record, func(record *kgo.Record, err error) {
		if err != nil {
			log.Printf("Produce failed: %v\n", err)
			return
		}
		log.Printf("Message sent to topic %s partition %d at offset %d\n", record.Topic, record.Partition, record.Offset)
	})

	k.client.Flush(context.Background())
}
