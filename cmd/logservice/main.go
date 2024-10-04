package main

import (
	"Resultify/log"
	"Resultify/registry"
	"Resultify/service"
	"context"
	"fmt"
	sllog "log"
)

func main() {
	log.Run("./Resultify.log")
	host, port := "localhost", "8080"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers)

	if err != nil {
		sllog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("log service is being shut down")
}
