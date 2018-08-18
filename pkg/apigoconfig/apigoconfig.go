package apigoconfig

import (
	"log"

	"github.com/kardianos/osext"
	"github.com/spf13/viper"
)

//Get ...
func Get() (*viper.Viper, error) {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	v := viper.New()

	v.SetDefault("logFolder", "${TEMP}")
	v.SetDefault("JobsDatabase.AdminDatabase", "postgres")
	v.SetDefault("JobsDatabase.Driver", "postgres")
	v.SetDefault("JobsDatabase.ConnectionString", "Host=localhost;Port=5432;Database=agdatabase;Username=aguser;Password=agpassword;SSL Mode=Require;Trust Server Certificate=true")
	v.SetDefault("Bindings", "0.0.0.0:1203")
	v.SetDefault("Debug", false)

	v.SetConfigName("default")
	v.SetConfigType("json")
	v.AddConfigPath("config")
	v.AddConfigPath(".")
	v.AddConfigPath(folderPath)
	v.AddConfigPath(folderPath + "/config")
	err = v.ReadInConfig()
	if err != nil {
		log.Printf("couldn't read the configuration from file: \n%v\n", err)
	}
	return v, nil
}
