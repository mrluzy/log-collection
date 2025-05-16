package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log-collection/conf"
	"log-collection/global"
	"log-collection/initconfig"
	"time"
)

func main() {
	conf.InitConfig()
	// 1.初始化kafka
	initconfig.Kafka()
	initconfig.Tail()

	//
	for {
		line, ok := <-global.Tailf.Lines
		if !ok {
			logrus.Info("tailfile is closed")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println(line.Text)
		msg := &sarama.ProducerMessage{}
		topic := viper.GetString("kafka.topic")
		msg.Topic = topic
		msg.Value = sarama.StringEncoder(line.Text)

		//
		global.MsgChan <- msg
	}
}
