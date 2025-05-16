package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
}
