package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jianhan/ms-sui-ideas/handler"
	"github.com/jianhan/ms-sui-ideas/proto/occupation"
	cfgreader "github.com/jianhan/pkg/configs"
	"github.com/micro/go-micro"
	"github.com/nats-io/go-nats-streaming"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENVIRONMENT")).Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	spew.Dump(viper.GetString("mongodb.host"))

	// init nats streaming
	sc, err := stan.Connect(viper.GetString("nats_streaming.cluster"), os.Getenv("nats_streaming.client_id"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	// get mongodb db
	session, err := mgo.Dial(viper.GetString("localhost"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// initialize new service
	srv := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)

	occupation.RegisterOccupationServiceHandler(
		srv.Server(),
		handler.NewOccupation(categoriesDB, productsDB, sc),
	)

	// init service
	srv.Init()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

func init() {
	// set default env if there is not one
	if os.Getenv("ENVIRONMENT") == "" {
		os.Setenv("ENVIRONMENT", "development")
	}

}
