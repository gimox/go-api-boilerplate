// package config provides  configuration by environment
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {

	// display info about config
	printInfo(env)

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)

	if env == "test" {
		v.AddConfigPath("../config/")
	} else {
		v.AddConfigPath("./config/")
	}

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	config = v
}

// main function call this to get config
func GetConfig() *viper.Viper {
	return config
}

// print some info on init
func printInfo(env string) {
	fmt.Println("")
	fmt.Println("***************************************************")
	fmt.Println("Application started with config:", env)
	fmt.Println("***************************************************")
	fmt.Println("")
}
