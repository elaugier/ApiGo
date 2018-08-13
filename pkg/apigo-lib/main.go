package main

//export dbConfig
type dbConfig struct {
	AdminDatabase    string `json:"AdminDatabase"`
	ConnectionString string `json:"ConnectionString"`
	Driver           string `json:"Driver"`
}

//export kafkaProducer
type kafkaProducer struct {
	BootstrapServers string `json:"Bootstrap.servers"`
}

//export configFileEngine
type configFileEngine struct {
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

func main() {}
