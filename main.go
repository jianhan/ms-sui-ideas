package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	// log .env
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
		panic(err)
	}

	// initialize new service
	srv := micro.NewService(
		micro.Name(viper.GetString("SERVICE_NAME")),
		micro.RegisterTTL(time.Duration(viper.GetInt("REGISTER_TTL_DURATION"))*time.Second),
		micro.RegisterInterval(time.Duration(viper.GetInt("REGISTER_INTERVAL"))*10),
		micro.Version(viper.GetString("SERVICE_VERSION")),
	)

	// init service
	srv.Init()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
