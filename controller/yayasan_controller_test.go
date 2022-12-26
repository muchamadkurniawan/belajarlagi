package controller

import (
	"belajarlagi/app"
	"belajarlagi/exception"
	"belajarlagi/repository"
	"belajarlagi/service"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestYayasanControllerImpl_FindAll(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepository := repository.NewYayasanRepository()
	yayasanService := service.NewYayasanService(yayasanRepository, db, validate)
	yayasanController := NewYayasanController(yayasanService)

	router := httprouter.New()
	router.GET("/api/yayasans", yayasanController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/yayasans", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	var yayasans = responseBody["data"].([]interface{})
	for index, yayasan := range yayasans {
		fmt.Println(yayasan)
		yayasanResponse := yayasans[index].(map[string]interface{})
		fmt.Println("yayasan id : ", yayasanResponse["id"])
		fmt.Println("yayasan nama : ", yayasanResponse["nama"])
		fmt.Println("yayasan uname : ", yayasanResponse["uname"])
		fmt.Println("yayasan pass : ", yayasanResponse["pass"])
	}
}
