package apigohandlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//SynchronousJob ...
func SynchronousJob(c *gin.Context) {
	value, exists := c.Get("id")
	str := fmt.Sprintf("%v", value)
	if exists {
		c.JSON(200, gin.H{
			"message": "synchronous job: OK" + str,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "synchronous job: OK",
		})
	}
}

//AsynchronousJob ...
func AsynchronousJob(c *gin.Context) {
	value, exists := c.Get("id")
	str := fmt.Sprintf("%v", value)
	if exists {
		c.JSON(200, gin.H{
			"message": "asynchronous job: OK" + str,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "asynchronous job: OK",
		})
	}
}
