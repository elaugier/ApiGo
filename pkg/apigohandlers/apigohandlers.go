package apigohandlers

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/restruct.v1"

	"github.com/gin-gonic/gin"
)

//SynchronousJob ...
func SynchronousJob(c *gin.Context) {
	buf, exists := c.Get("route")
	if exists {
		var cr viper.Viper
		restruct.Unpack(buf.([]byte), binary.LittleEndian, &cr)
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("%v", cr),
		})
	} else {
		c.JSON(500, gin.H{
			"msg": "error : no route key in current context",
		})
	}
}

//AsynchronousJob ...
func AsynchronousJob(c *gin.Context) {
	buf, exists := c.Get("route")
	if exists {
		var cr viper.Viper
		restruct.Unpack(buf.([]byte), binary.LittleEndian, &cr)
		for i, v := range cr.AllSettings() {
			log.Printf("%s : %v", i, v)
		}
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("%v", cr),
		})
	} else {
		c.JSON(500, gin.H{
			"msg": "error : no route key in current context",
		})
	}
}

func getRouteKey(key string, c *gin.Context) string {
	value, exists := c.Get(key)
	if exists {
		return fmt.Sprintf("%v", value)
	}
	return ""
}
