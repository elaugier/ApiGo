package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/elaugier/ApiGo/pkg/apigolib"
)

func main() {

	// args := os.Args

	f, err := os.OpenFile("apigo.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	jsonFile, err := os.Open("config/default.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully Opened config/default.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config apigolib.ConfigFileEngine
	json.Unmarshal(byteValue, &config)

}
