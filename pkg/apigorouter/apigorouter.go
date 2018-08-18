package apigorouter

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/elaugier/ApiGo/pkg/apigohandlers"

	"github.com/elaugier/ApiGo/pkg/apigoconfig"
	"github.com/spf13/viper"

	"github.com/bmatcuk/doublestar"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/osext"
)

//CurrentRoute ...
func CurrentRoute(id int, route *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("id", id)
		c.Set("route", route)
		c.Next()
	}
}

//Get ...
func Get(pathConfig string) (*gin.Engine, error) {

	log.Println("Create default gin engine")
	r := gin.Default()

	log.Println("setup '/ping' route")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	pattern := pathConfig + "/**/*.conf.json"
	volumeName := filepath.VolumeName(pattern)
	if volumeName == "" || strings.HasPrefix(pathConfig, "/") {
		pattern = folderPath + "/" + pathConfig + "/**/*.conf.json"
	}
	log.Printf("Try to retrieve routes configurations in path : %s", pattern)

	filesConf, err := doublestar.Glob(pattern)
	if err != nil {

		log.Printf("error on recursive search for *.conf.json in folder : %s => %v", pathConfig, err)
	}

	routesConfigs := make(map[int]*viper.Viper)
	for i, f := range filesConf {
		routesConfigs[i] = apigoconfig.GetRouteConfig(f)
		log.Printf("(%d ==> %s", i, f)
	}

	for index, route := range routesConfigs {
		routeName := route.GetString("Name")
		routePath := strings.ToLower(route.GetString("Route"))
		jobType := strings.ToLower(route.GetString("JobType"))
		method := strings.ToLower(route.GetString("Method"))
		switch jobType {
		case "synchronous":

			switch method {
			case "get":
				r.GET(routePath, CurrentRoute(index, route), apigohandlers.SynchronousJob)

			case "post":
				r.POST(routePath, CurrentRoute(index, route), apigohandlers.SynchronousJob)

			case "put":
				r.PUT(routePath, CurrentRoute(index, route), apigohandlers.SynchronousJob)

			case "patch":
				r.PATCH(routePath, CurrentRoute(index, route), apigohandlers.SynchronousJob)

			case "delete":
				r.DELETE(routePath, CurrentRoute(index, route), apigohandlers.SynchronousJob)

			default:
				log.Printf("Unknown method or invalid method for route '%s'", routeName)
			}

		case "asynchronous":
			switch method {
			case "get":
				r.GET(routePath, CurrentRoute(index, route), apigohandlers.AsynchronousJob)

			case "post":
				r.POST(routePath, CurrentRoute(index, route), apigohandlers.AsynchronousJob)

			case "put":
				r.PUT(routePath, CurrentRoute(index, route), apigohandlers.AsynchronousJob)

			case "patch":
				r.PATCH(routePath, CurrentRoute(index, route), apigohandlers.AsynchronousJob)

			case "delete":
				r.DELETE(routePath, CurrentRoute(index, route), apigohandlers.AsynchronousJob)

			default:
				log.Printf("Unknown method or invalid method for route '%s'", routeName)
			}
		default:
			log.Printf("Unknown job type for route '%s'", routeName)
		}
	}

	return r, nil
}
