package main

import (
	"fmt"

	scrennshotapp "github.com/aravion1/Scrennshoter"
	"github.com/aravion1/Scrennshoter/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("MAIN OK")

	if err := InitConfig(); err != nil {
		panic(err.Error())
	}

	h := handler.NewHandler()
	s := new(scrennshotapp.Server)
	if err := s.Run("4000", h.GetHandlers()); err != nil {
		logrus.Fatal("Run error " + err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
