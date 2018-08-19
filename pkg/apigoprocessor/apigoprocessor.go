package apigoprocessor

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/elaugier/ApiGo/pkg/apigolib"
)

//Process ...
func Process(message *kafka.Message) {
	var msg apigolib.JsonCmd
	r := bytes.NewReader(message.Value)
	jsonParser := json.NewDecoder(r)
	jsonParser.Decode(&msg)
	switch msg.Type {
	case "Powershell":
		log.Printf("Powershell type detected (%s)\r\n", msg.Uuid)
		//TODO: write PowershellHelper
	case "Python":
		log.Printf("Python type detected (%s)\r\n", msg.Uuid)
		//TODO: write PythonHelper
	case "Perl":
		log.Printf("Perl type detected (%s)\r\n", msg.Uuid)
		//TODO: write PerlHelper
	case "CommandLine":
		log.Printf("CLI type detected (%s)\r\n", msg.Uuid)
		//TODO: write CLIHelper
	default:
		log.Printf("Unknown type detected for job %s\r\n", msg.Uuid)
	}

	//TODO: Update job state in database
}
