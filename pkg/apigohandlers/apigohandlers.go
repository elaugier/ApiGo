package apigohandlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/elaugier/ApiGo/pkg/apigolib"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

//RoutesConfigs ...
var RoutesConfigs map[int]*viper.Viper

//SynchronousJob ...
func SynchronousJob(c *gin.Context) {
	apigolib.Trace()
	buf, _ := c.Get("id")
	id := buf.(int)
	var Route RouteConfig
	err := RoutesConfigs[id].Unmarshal(&Route)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	currentRoute := Route.Name

	for i := 0; i < len(Route.Cmd.Params); i++ {

		p := Route.Cmd.Params[i]

		log.Printf("Expected parameter name: %s", p.Name)

		switch p.In {
		case "uri":
			value := c.Param(p.Name)
			log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

		case "header":
			value := c.GetHeader(p.Name)
			log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

		case "querystring":
			value := c.Query(p.Name)
			log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

		case "body":
			var keyValue map[string]string
			c.BindJSON(&keyValue)
			value := keyValue[p.Name]
			log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

		default:
			log.Printf("Unkown 'In' value for param '%s'", p.Name)

		}
	}

	log.Printf("Current Route => %s", currentRoute)
	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("%s", currentRoute),
	})
}

//AsynchronousJob ...
func AsynchronousJob(c *gin.Context) {
	apigolib.Trace()
	buf, _ := c.Get("id")
	id := buf.(int)
	var Route RouteConfig
	err := RoutesConfigs[id].Unmarshal(&Route)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	currentRoute := Route.Name

	for i := 0; i < len(Route.Cmd.Params); i++ {

		p := Route.Cmd.Params[i]

		log.Printf("Expected parameter name: %s", p.Name)
		mandatory, err := strconv.ParseBool(p.Mandatory)
		if err == nil {
			switch p.In {
			case "uri":
				value := c.Param(p.Name)
				log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

				if value == "" && mandatory {
					log.Printf("Parameter '%s' is mandatory => raise error and add message for response", p.Name)
				}
				if value == "" && !mandatory {
					log.Printf("Parameter '%s' is not mandatory but empty => no action", p.Name)
				}

			case "header":
				value := c.GetHeader(p.Name)
				log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

			case "querystring":
				value := c.Query(p.Name)
				log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

			case "body":
				var keyValue map[string]string
				c.BindJSON(&keyValue)
				value := keyValue[p.Name]
				log.Printf("retrieve key '%s' => '%s' from %s", p.Name, value, p.In)

			default:
				log.Printf("Unkown 'In' value for param '%s'", p.Name)

			}
		}
	}

	log.Printf("Current Route => %s", currentRoute)
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
