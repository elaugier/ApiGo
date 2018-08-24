package apigohelpers

import (
	"encoding/json"
	"log"

	"github.com/elaugier/ApiGo/pkg/apigoconfig"
	"github.com/elaugier/ApiGo/pkg/apigokafka"
	"github.com/spf13/viper"
)

//NewKafka ...
func NewKafka() *Kafka {
	config, err := apigoconfig.Get()
	if err != nil {
		log.Fatalf("Error on get configuration file : %v", err)
	}
	return &Kafka{c: *config}
}

//Kafka ...
type Kafka struct {
	P apigokafka.Producer
	c viper.Viper
}

//Send ...
func (k Kafka) Send(j JSONCmd, topic string) error {
	buf, err := json.Marshal(&j)
	if err != nil {
		log.Printf("Error on marshaling JSONCmd struct : %v", err)
		return err
	}
	s := string(buf)
	err = k.P.Send(s, topic)
	if err != nil {
		log.Printf("Error on sending JSON to Kafka producer")
		return err
	}
	log.Printf("JSONCmd sent to Kafka producer")
	return nil
}

//JSONCmd ...
type JSONCmd struct {
	UUID     string            `json:"Uuid"`
	Name     string            `json:"Name"`
	Type     string            `json:"Type"`
	PSModule string            `json:"PSModule"`
	PyVenv   string            `json:"PyVenv"`
	Params   map[string]string `json:"Params"`
	JobType  string            `json:"JobType"`
	Timeout  string            `json:"Timeout"`
}
