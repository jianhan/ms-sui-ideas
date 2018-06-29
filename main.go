package main

import (
	"time"

	"github.com/jianhan/ms-sui-ideas/handler"
	"github.com/micro/go-micro"
	_ "github.com/spf13/viper/remote"
)

func main() {
	// initialize new service
	srv := micro.NewService(
		micro.Name("Ideas"),
		micro.RegisterTTL(time.Duration(50)*time.Second),
		micro.RegisterInterval(time.Duration(30)*10),
		micro.Version("1.0.0"),
	)
	handler.NewOccupation(nil, nil)
	// init service
	srv.Init()
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
