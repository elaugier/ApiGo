package apigoprocessor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/elaugier/ApiGo/pkg/apigohelpers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/elaugier/ApiGo/pkg/apigolib"
)

//Process ...
func Process(message *kafka.Message) {
	var msg apigolib.JsonCmd
	exitCode := 0
	result := ""
	var err error
	r := bytes.NewReader(message.Value)
	jsonParser := json.NewDecoder(r)
	jsonParser.Decode(&msg)
	switch msg.Type {
	case "Powershell":
		log.Printf("Powershell type detected (%s)\r\n", msg.Uuid)
		exitCode, result, err = apigohelpers.PowershellRun(msg.PSModule, msg.Name, msg.Args)
	case "Python":
		log.Printf("Python type detected (%s)\r\n", msg.Uuid)
		exitCode, result, err = apigohelpers.PythonRun(msg.PyEnv, msg.Name, msg.Args)
	case "Perl":
		log.Printf("Perl type detected (%s)\r\n", msg.Uuid)
		exitCode, result, err = apigohelpers.PerlRun(msg.Name, msg.Args)
	case "Ruby":
		log.Printf("Ruby type detected (%s)\r\n", msg.Uuid)
		exitCode, result, err = apigohelpers.RubyRun(msg.Name, msg.Args)
	case "CommandLine":
		log.Printf("CLI type detected (%s)\r\n", msg.Uuid)
		exitCode, result, err = apigohelpers.CLIRun(msg.Name, msg.Args)
	default:
		log.Printf("Unknown type detected for job %s\r\n", msg.Uuid)
		exitCode = 1
		result = fmt.Sprintf("Unknown type detected for job %s : %s", msg.Uuid, msg.Type)
		err = errors.New(result)
	}

	if exitCode != 0 || err != nil {

	} else {

	}
	//TODO: Update job state in database
}
