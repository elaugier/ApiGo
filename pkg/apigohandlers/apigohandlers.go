package apigohandlers

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

//RoutesConfigs ...
var RoutesConfigs map[int]*viper.Viper

//SynchronousJob ...
func SynchronousJob(c *gin.Context) {
	buf, _ := c.Get("id")
	id := buf.(int)

	currentRoute := RoutesConfigs[id].GetString("Name")

	expectedParams := RoutesConfigs[id].GetStringMapString("Cmd.Params")

	for p, v := range expectedParams {
		log.Printf("param %s => %s", p, v)
	}
	//request := c.Request

	//p := apigokafka.Producer

	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("%s", currentRoute),
	})
}

//AsynchronousJob ...
func AsynchronousJob(c *gin.Context) {
	buf, _ := c.Get("id")
	id := buf.(int)

	currentRoute := RoutesConfigs[id].GetString("Name")

	expectedParams := RoutesConfigs[id].GetStringMapString("Cmd.Params")

	for p, v := range expectedParams {
		log.Printf("param %s => %s", p, v)
	}

	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("%s", currentRoute),
	})
}

//Parameter ...
type Parameter struct {
	Name      string `json:"Name"`
	Type      string `json:"Type"`
	Mandatory string `json:"Mandatory"`
	In        string `json:"In"`
}

//Cmd ...
type Cmd struct {
	Name     string      `json:"Name"`
	Type     string      `json:"Type"`
	PSModule string      `json:"PSModule"`
	PyVenv   string      `json:"PyVenv"`
	Params   []Parameter `json:"Params"`
}

//RouteConfig ...
type RouteConfig struct {
	Name              string `json:"Name"`
	Cmd               Cmd    `json:"Cmd"`
	Route             string `json:"Route"`
	Method            string `json:"Method"`
	JobType           string `json:"JobType"`
	Topic             string `json:"Topic"`
	Timeout           string `json:"Timeout"`
	AddRequestIDParam string `json:"AddRequestIDParam"`
}
