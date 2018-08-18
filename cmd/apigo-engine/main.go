package main

import (
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/osext"

	"github.com/elaugier/ApiGo/pkg/apigoconfig"

	"github.com/gin-gonic/gin"
)

/* func logConfig(config map[string]interface, sep string){
	for key, value := range config {
		typeValue := reflect.TypeOf(value)
		if(typeValue=="map[string]"){
			sep += " "
			logConfig(value, sep)
		}
		else {
			log.Printf("%s => %s", key, value)
		}
	}
} */

func main() {

	fullBinaryName, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	binaryName := strings.Replace(strings.Replace(fullBinaryName, folderPath, "", -1), string(os.PathSeparator), "", -1)

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	log.SetPrefix(binaryName + " " + strconv.Itoa(os.Getpid()) + " ")

	config, err := apigoconfig.Get()
	if err != nil {
		log.Fatal(err)
	}

	timestampStart := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	logFile := os.ExpandEnv(config.GetString("logFolder")) + "/" + timestampStart + "_" + binaryName + ".log"
	log.Println("log file location => '" + logFile + "'")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	// TODO: Output some parameters from read config
	log.Printf("typeof => %s", reflect.TypeOf(config.AllSettings()))

	if config.GetBool("Debug") {
		log.Println("enable gin TestMode")
		gin.SetMode(gin.TestMode)
	} else {
		log.Println("enable gin ReleaseMode")
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.GetString("Bindings"))
}
