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
	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("%s", RoutesConfigs[id].GetString("Name")),
	})
}
