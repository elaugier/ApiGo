package apigoconfig

import (
	"log"

	"github.com/spf13/viper"
)

//Get ...
func Get() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't read the configuration: %v\n", err)
		return nil, err
	}
	return v, nil
}
