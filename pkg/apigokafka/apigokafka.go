package apigokafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

//Producer ...
var Producer producer

type producer struct {
	P *kafka.Producer
}

//InitConnection ...
func (p producer) InitConnection(config viper.Viper) {
	var err error
	p.P, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Printf("Kafka connection failed : %v", err)
	}
	return
}

//Send ...
func (producer) Send(msg string) {

}
