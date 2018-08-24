package apigomiddleware

import (
	"log"

	"github.com/elaugier/ApiGo/pkg/apigolib"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//RequestUUID ...
func RequestUUID() gin.HandlerFunc {
	apigolib.Trace()
	return func(c *gin.Context) {
		apigolib.Trace()
		requestUUID, err := uuid.NewUUID()
		if err != nil {
			log.Fatalf("Error on generate request UUID : %v", err)
		}
		c.Writer.Header().Set("X-Request-Id", requestUUID.String())
		log.Printf("set X-Request-id => %s", requestUUID.String())
		c.Next()
	}
}
