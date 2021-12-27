package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ReadConf(configFileLocation, configFileName string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(configFileLocation)

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading conf file %v\n", err)
		os.Exit(1)
	}

	return v
}
