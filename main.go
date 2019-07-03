package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/rocknio/micro-assets/assets-srv/handler"
	//"github.com/rocknio/micro-assets/assets-srv/subscriber"

	asset "github.com/rocknio/micro-assets/assets-srv/proto/asset"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("rz.assets.srv.asset"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = asset.RegisterAssetMethodHandler(service.Server(), new(handler.Asset))

	//// Register Struct as Subscriber
	//micro.RegisterSubscriber("rz.assets.srv.asset", service.Server(), new(subscriber.Asset))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("rz.assets.srv.asset", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
