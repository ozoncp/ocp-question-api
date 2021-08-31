package producer

import (
	"github.com/ozoncp/ocp-question-api/internal/config"
	"time"

	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

// Producer - an interface for send event messages
type Producer interface {
	Send(message Message) error
}

type producer struct {
	producer sarama.SyncProducer
	topic    string
}

// NewProducer - creates a new instance of Producer
func NewProducer() *producer {
	conf := config.NewConfig()

	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForLocal
	producerConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	producerConfig.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(conf.Kafka.Brokers, producerConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Sarama new sync producer")
	}

	return &producer{
		producer: syncProducer,
		topic:    conf.Kafka.Topic,
	}
}

func (p *producer) Send(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("failed marshaling message to json:")
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(p.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.producer.SendMessage(msg)
	return err
}
