package apigohandlers

import "github.com/gin-gonic/gin"

//SynchronousJob ...
func SynchronousJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "synchronous job: OK",
	})
}

//AsynchronousJob ...
func AsynchronousJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "asynchronous job: OK",
	})
}
