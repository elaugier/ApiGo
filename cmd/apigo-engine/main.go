package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type dbConfig struct {
	AdminDatabase    string `json:"AdminDatabase"`
	ConnectionString string `json:"ConnectionString"`
	Driver           string `json:"Driver"`
}

type kafkaProducer struct {
	BootstrapServers string `json:"Bootstrap.servers"`
}

type configFile struct {
	AccountingDatabase              dbConfig      `json:"AccountingDatabase"`
	Bindings                        string        `json:"Bindings"`
	CertPath                        string        `json:"CertPath"`
	CertPwd                         string        `json:"CertPwd"`
	JobsDatabase                    dbConfig      `json:"JobsDatabase"`
	KafkaProducer                   kafkaProducer `json:"KafkaProducer"`
	MaxConcurrentConnections        int64         `json:"MaxConcurrentConnections"`
	MaxConcurrentUpgradeConnections int64         `json:"MaxConcurrentUpgradeConnections"`
	MaxRequestBodySize              int64         `json:"MaxRequestBodySize"`
	RoutesConfigPath                string        `json:"RoutesConfigPath"`
	Secure                          bool          `json:"Secure"`
}

func main() {
	jsonFile, err := os.Open("config/default.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config/default.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config configFile
	json.Unmarshal(byteValue, &config)

}
