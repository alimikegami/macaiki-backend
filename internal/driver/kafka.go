package driver

import (
	"context"

	"github.com/segmentio/kafka-go"
)

var KafkaConn *kafka.Conn

func CreateKafkaProducer(host string, topic string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", host, topic, 0)
	if err != nil {
		panic(err)
	}

	KafkaConn = conn
}
