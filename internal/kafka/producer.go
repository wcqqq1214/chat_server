package kafka

import (
	"chat-room/pkg/global/log"
	"strings"

	"github.com/IBM/sarama"
)

var producer sarama.AsyncProducer
var topic string = "default_message"

func InitProducer(topicInput, hosts string) {
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Logger.Error("init kafka client error", log.Any("init kafka client error", err.Error()))
	}

	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Logger.Error("init kafka async client error", log.Any("init kafka async client error", err.Error()))
	}
}

func Send(data []byte) {
	be := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: be}
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}
