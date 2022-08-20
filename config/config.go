package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	//path, err := os.Getwd()
	//if err != nil {
	//	//	panic(err)
	//	//}
	viper.AddConfigPath("/jcoin/cake-syrup-pools-lover/config")
	//viper.AddConfigPath("D:\\Code\\github\\cake-syrup-pools-lover\\config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
