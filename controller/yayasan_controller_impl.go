package controller

import (
	"belajarlagi/helper"
	"belajarlagi/model/web"
	"belajarlagi/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
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
	//fmt.Println("ini controller FindById yayasan")

	//fmt.Println("opppo", request.URL.Query().Get("id"))
	//id := request.URL.Query().Get("id")
	id := params.ByName("id")
	//fmt.Println("idddd", id)
	YayasanId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	yayasanResponse := yayasanController.YayasanService.FindById(request.Context(), YayasanId)
	fmt.Println("yayasan response", yayasanResponse)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   yayasanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
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
