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

//NewProcessor ...
func NewProcessor() Processor {
	return Processor{
		currentJobsCount: 0,
	}
}

//Processor ...
type Processor struct {
	currentJobsCount int64
}

//GetCurrentJobsCount ...
func (p Processor) GetCurrentJobsCount() int64 {
	return p.currentJobsCount
}

//Process ...
func (p Processor) Process(message *kafka.Message, done chan string) {
	var msg apigolib.JSONCmd
	p.currentJobsCount++
	exitCode := 0
	result := ""
	var err error
	r := bytes.NewReader(message.Value)
	jsonParser := json.NewDecoder(r)
	jsonParser.Decode(&msg)
	switch msg.Type {
	case "Powershell":
		log.Printf("Powershell type detected (%s)\r\n", msg.UUID)
		exitCode, result, err = apigohelpers.PowershellRun(msg.PSModule, msg.Name, msg.Args)
	case "Python":
		log.Printf("Python type detected (%s)\r\n", msg.UUID)
		exitCode, result, err = apigohelpers.PythonRun(msg.PyEnv, msg.Name, msg.Args)
	case "Perl":
		log.Printf("Perl type detected (%s)\r\n", msg.UUID)
		exitCode, result, err = apigohelpers.PerlRun(msg.Name, msg.Args)
	case "Ruby":
		log.Printf("Ruby type detected (%s)\r\n", msg.UUID)
		exitCode, result, err = apigohelpers.RubyRun(msg.Name, msg.Args)
	case "CommandLine":
		log.Printf("CLI type detected (%s)\r\n", msg.UUID)
		exitCode, result, err = apigohelpers.CLIRun(msg.Name, msg.Args)
	default:
		log.Printf("Unknown type detected for job %s\r\n", msg.UUID)
		exitCode = 1
		result = fmt.Sprintf("Unknown type detected for job %s : %s", msg.UUID, msg.Type)
		err = errors.New(result)
	}
	p.currentJobsCount--
	done <- "OK"
	if exitCode != 0 || err != nil {

	} else {

	}
	//TODO: Update job state in database
}
