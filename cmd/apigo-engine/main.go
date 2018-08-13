package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elaugier/ApiGo/pkg/apigolib"
)

func main() {
	jsonFile, err := os.Open("config/default.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config/default.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config apigolib.ConfigFileEngine
	json.Unmarshal(byteValue, &config)

}
