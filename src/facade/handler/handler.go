package handler

import (
	"syscommon"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gitstliu/go-id-worker"
	"github.com/gitstliu/log4go"
)

type IDWorkerFacade struct {
}

var currWoker = &idworker.IdWorker{}

func (this *IDWorkerFacade) InitIdWorker(workerId, datacenterId int64) error {
	return currWoker.InitIdWorker(workerId, datacenterId)
}

func (this *IDWorkerFacade) GetNewID(w rest.ResponseWriter, r *rest.Request) {
	response := syscommon.CommonResponse{Code: syscommon.Success, Message: "Success"}
	id, nextIDErr := currWoker.NextId()

	if nextIDErr != nil {
		log4go.Error(nextIDErr)
		response.Code = syscommon.Fail
		response.Message = "Get ID Error, Please Check!!"
		w.WriteJson(&response)
		return
	}
	response.Result = id
	log4go.Debug("CurrID is %v", id)
	w.WriteJson(response)
	return
}
