package apigomiddleware

import (
	"log"

	"github.com/elaugier/ApiGo/pkg/apigolib"
	"github.com/gin-gonic/gin"
)

//Db ...
func Db() gin.HandlerFunc {
	apigolib.Trace()
	log.Println("create DB connection")
	toto := "OK"
	return func(c *gin.Context) {
		apigolib.Trace()
		log.Printf("use DB connection (toto = %s)", toto)
		c.Next()
	}
}
