package apigomiddleware

import (
	"github.com/elaugier/ApiGo/pkg/apigolib"
	"github.com/gin-gonic/gin"
)

//Apikey ...
func Apikey(c *gin.Context) {
	apigolib.Trace()
	c.Next()
}
