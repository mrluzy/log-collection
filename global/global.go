package global

import (
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
)

var (
	KafkaClient sarama.SyncProducer
	Tailf       *tail.Tail
	MsgChan     chan *sarama.ProducerMessage
)
