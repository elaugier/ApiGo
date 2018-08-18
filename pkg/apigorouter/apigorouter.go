package apigorouter

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/elaugier/ApiGo/pkg/apigoconfig"
	"github.com/spf13/viper"

	"github.com/bmatcuk/doublestar"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/osext"
)

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

	return r, nil
}
