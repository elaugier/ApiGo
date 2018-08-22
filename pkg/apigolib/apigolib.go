package apigolib

import (
	"log"
	"runtime"
)

//DbConfig ...
type DbConfig struct {
	AdminDatabase    string `json:"AdminDatabase"`
	ConnectionString string `json:"ConnectionString"`
	Driver           string `json:"Driver"`
}

//KafkaProducer ...
type KafkaProducer struct {
	BootstrapServers string `json:"Bootstrap.servers"`
}

//ConfigFileEngine ...
type ConfigFileEngine struct {
	AccountingDatabase              DbConfig      `json:"AccountingDatabase"`
	Bindings                        string        `json:"Bindings"`
	CertPath                        string        `json:"CertPath"`
	CertPwd                         string        `json:"CertPwd"`
	JobsDatabase                    DbConfig      `json:"JobsDatabase"`
	KafkaProducer                   KafkaProducer `json:"KafkaProducer"`
	MaxConcurrentConnections        int64         `json:"MaxConcurrentConnections"`
	MaxConcurrentUpgradeConnections int64         `json:"MaxConcurrentUpgradeConnections"`
	MaxRequestBodySize              int64         `json:"MaxRequestBodySize"`
	RoutesConfigPath                string        `json:"RoutesConfigPath"`
	Secure                          bool          `json:"Secure"`
}

//JSONCmd ...
type JSONCmd struct {
	UUID     string            `json:"Uuid"`
	Name     string            `json:"Name"`
	Type     string            `json:"Type"`
	PSModule string            `json:"PSModule"`
	PyEnv    string            `json:"PyEnv"`
	Args     map[string]string `json:"Args"`
}

//Trace ...
func Trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	log.Printf("Entering in %s\n", f.Name())
}
