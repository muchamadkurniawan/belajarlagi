package controller

import (
	"belajarlagi/helper"
	"belajarlagi/model/web"
	"belajarlagi/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type YayasanControllerImpl struct {
	YayasanService service.YayasanService
}

func (yayasanController *YayasanControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (yayasanController *YayasanControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (yayasanController *YayasanControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (yayasanController *YayasanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (yayasanController *YayasanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responses := yayasanController.YayasanService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func NewYayasanController(yayasanService service.YayasanService) YayasanController {
	return &YayasanControllerImpl{YayasanService: yayasanService}
}
