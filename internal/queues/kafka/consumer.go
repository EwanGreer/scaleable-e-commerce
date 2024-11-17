package kafka

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer interface {
	Consume() error
}

type KafkaConsumer struct {
	brokers []string
	group   string
	topics  []string
}

func NewConsumer(brokers []string, consumerGroup string, topics []string) *KafkaConsumer {
	return &KafkaConsumer{
		brokers: brokers,
		group:   consumerGroup,
		topics:  topics,
	}
}

func (c KafkaConsumer) Consume() error {
	opts := []kgo.Opt{
		kgo.SeedBrokers(c.brokers[0]),
		kgo.ConsumerGroup(c.group),
		kgo.ConsumeTopics(c.topics[0]),
		kgo.AutoCommitMarks(),
		kgo.AutoCommitInterval(5 * time.Second),
		kgo.HeartbeatInterval(3 * time.Second),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		slog.Error("Error creating Kafka client", "err", err)
		os.Exit(1)
	}
	defer client.Close()

	ctx := context.Background()

	for {
		slog.Info("Polling...", "topics", c.topics)
		fetches := client.PollFetches(ctx)

		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				slog.Error("Error consuming messages",
					"topic", err.Topic,
					"partition", err.Partition,
					"error", err.Err,
				)
			}
			continue
		}

		records := fetches.Records()
		for _, record := range records {
			slog.Info("Received message",
				"topic", record.Topic,
				"partition", record.Partition,
				"offset", record.Offset,
				"key", string(record.Key),
				"value", string(record.Value),
			)

			client.MarkCommitRecords(records...)
		}
	}
}
