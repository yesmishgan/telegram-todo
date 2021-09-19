package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		zap.NewPro
	}
	fmt.Println("Test")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
