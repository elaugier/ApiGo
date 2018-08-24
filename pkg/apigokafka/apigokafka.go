package apigokafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

//NewProducer ...
func NewProducer(config viper.Viper) *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.GetString("KafkaProducer.BootstrapServers")})
	if err != nil {
		log.Printf("Kafka connection failed : %v", err)
	}
	return &Producer{P: p}
}

//Producer ...
type Producer struct {
	P *kafka.Producer
}

//CloseConnection ...
func (p Producer) CloseConnection() {
	p.P.Close()
	return
}

//Send ...
func (p Producer) Send(msg string, topic string) error {
	err := p.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, nil)
	if err != nil {
		log.Printf("Error on produce message to Kafka : %v", err)
		return err
	}
	log.Printf("Message sent to Kafka service")
	p.P.Flush(10000)
	return nil
}
