package kafka

import (
	"chat-room/pkg/global/log"
	"strings"

	"github.com/IBM/sarama"
)

var consumer sarama.Consumer

type ConsumerCallback func(data []byte)

// 初始化消费者
func InitConsumer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Logger.Error("init kafka consumer client error", log.Any("init kafka consumer client error", err.Error()))
	}

	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Logger.Error("init kafka consumer error", log.Any("init kafka consumer error", err.Error()))
	}
}

// 消费信息，通过回调函数进行
func ConsumerMsg(callBack ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Logger.Error("ConsumePartition error", log.Any("ConsumePartition error", err.Error()))
		return
	}

	defer partitionConsumer.Close()
	for {
		msg := <-partitionConsumer.Messages()
		if callBack != nil {
			callBack(msg.Value)
		}
	}
}

func CloseConsumer() {
	if consumer != nil {
		consumer.Close()
	}
}
