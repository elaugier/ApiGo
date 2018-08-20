package apigohandlers

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

//RoutesConfigs ...
var RoutesConfigs map[int]*viper.Viper

//SynchronousJob ...
func SynchronousJob(c *gin.Context) {
	buf, _ := c.Get("id")
	id := buf.(int)

	//p := apigokafka.Producer

	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("%s", RoutesConfigs[id].GetString("Name")),
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
