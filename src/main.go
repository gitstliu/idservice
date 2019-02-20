package main

import (
	"config"
	"facade/handler"
	"fmt"
	"web/restadapter"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gitstliu/log4go"
)

func main() {

	defer panicHandler()

	log4go.LoadConfiguration("config/log.xml")
	defer log4go.Close()

	config.LoadConfigure("config/config.toml")

	log4go.Info("Config Load Success!!!")

	restAdapter := restadapter.RestAdapter{initUrl(), config.GetConfigure().ServicePort}

	log4go.Info("Starting RestAdapter!!!")
	restAdapter.Start()
}

func initUrl() []*restadapter.UrlMap {
	idWorkerFacade := handler.IDWorkerFacade{}
	idWorkerFacade.InitIdWorker(config.GetConfigure().WorkerId, config.GetConfigure().DatacenterId)

	urls := make([]*restadapter.UrlMap, 0, 100)
	urls = append(urls, &restadapter.UrlMap{
		Url: "/ids",
		MethodMap: map[string]rest.HandlerFunc{
			"POST": idWorkerFacade.GetNewID}})

	return urls
}

func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Printf("%T", r)
		panic(r)
	}
}
