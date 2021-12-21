package configuration

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// ReadConf loads the .confs file to Configuration
// TODO: Generalize these paths, if we want tests in the future the modules that need cfg will fail
// Or each test will need to copy this function
func ReadConf() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs/")
	v.AddConfigPath("./configuration/")
	v.SetEnvPrefix("SSP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println("Error reading conf file")
		os.Exit(1)
	}

	return v
}
