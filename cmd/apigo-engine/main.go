package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/osext"

	"github.com/elaugier/ApiGo/pkg/apigoconfig"
	"github.com/elaugier/ApiGo/pkg/apigorouter"

	"github.com/gin-gonic/gin"
)

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

	if config.GetBool("Debug") {
		log.Println("enable gin DebugMode")
		gin.SetMode(gin.DebugMode)
	} else {
		log.Println("enable gin ReleaseMode")
		gin.SetMode(gin.ReleaseMode)
	}

	r, err := apigorouter.Get(config.GetString("RoutesConfigPath"))
	if err != nil {
		log.Panicf("Loading routes : FAILED! => %v", err)
	}

	r.Run(config.GetString("Bindings"))
}
