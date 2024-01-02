package commons

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	os.Mkdir("config", 0666)
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
}

func GetConfig(key string) interface{} {
	return viper.Get(key)
}

func SetConfig(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
}
