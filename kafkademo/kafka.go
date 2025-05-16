package kafkademo

import (
	"fmt"
	"github.com/IBM/sarama"
)

func connectKafka() {
	// 1.producer configuration
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.Return.Successes = true

	// 2.connect to Kafka
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 3. 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("hello world")

	// 4.send message
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
