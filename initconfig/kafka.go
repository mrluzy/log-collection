package initconfig

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log-collection/global"
)

func Kafka() {
	// 初始化配置文件
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.Return.Successes = true

	// 连接Kafka
	addr := viper.GetString("kafka.addr")
	client, err := sarama.NewSyncProducer([]string{addr}, cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	//defer client.Close()
	global.KafkaClient = client
	global.MsgChan = make(chan *sarama.ProducerMessage, 10000)
	logrus.Info("kafka init success")
	go sendMessage()
}

func sendMessage() {
	for {
		select {
		case msg := <-global.MsgChan:
			pid, offset, err := global.KafkaClient.SendMessage(msg)
			if err != nil {
				logrus.Warnf("send message error", err)
				return
			}
			logrus.Infof("send message to kafka success, pid=%d, offset=%d", pid, offset)
		}
	}
}
