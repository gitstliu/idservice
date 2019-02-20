package restadapter

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gitstliu/log4go"
)

type RequestHandler func(w rest.ResponseWriter, r *rest.Request)

type RestAdapter struct {
	UrlList []*UrlMap
	Port    int
}

func (restAdapter *RestAdapter) SetPort(port int) {
	restAdapter.Port = port
}

func (restAdapter *RestAdapter) GetPort() int {
	return restAdapter.Port
}

func (restAdapter *RestAdapter) Start() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	routesRule := make([]*rest.Route, 0, 100)

	for _, v := range restAdapter.UrlList {
		routesRule = append(routesRule, v.ToRoute()...)
	}

	router, err := rest.MakeRouter(routesRule...)

	if err != nil {
		log4go.Error(err)
	}
	api.SetApp(router)
	log4go.Info("Start Port %v", restAdapter.Port)
	log4go.Error(http.ListenAndServe(":"+strconv.Itoa(restAdapter.Port), api.MakeHandler()))
}

type UrlMap struct {
	Url string
	//MethodMap map[string]func(w rest.ResponseWriter, r *rest.Request)
	MethodMap map[string]rest.HandlerFunc
}

func (urlMap *UrlMap) ToRoute() []*rest.Route {

	result := make([]*rest.Route, 0, 5)
	for k, v := range urlMap.MethodMap {

		if strings.EqualFold(k, "GET") {
			result = append(result, rest.Get(urlMap.Url, v))
		} else if strings.EqualFold(k, "POST") {
			result = append(result, rest.Post(urlMap.Url, v))
		} else if strings.EqualFold(k, "PUT") {
			result = append(result, rest.Put(urlMap.Url, v))
		} else if strings.EqualFold(k, "DELETE") {
			result = append(result, rest.Delete(urlMap.Url, v))
		} else if strings.EqualFold(k, "PATCH") {
			result = append(result, rest.Patch(urlMap.Url, v))
		}
	}

	return result
}
