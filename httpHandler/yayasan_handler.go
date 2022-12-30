package httpHandler

import (
	"belajarlagi/app"
	"belajarlagi/controller"
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
	"path"
	"text/template"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index2.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := app.NewDB()
	//tx, _ := db.Begin()
	validate := validator.New()
	yayasanRepo := repository.NewYayasanRepository()
	yayasanService := service.NewYayasanService(yayasanRepo, db, validate)
	yayasanController := controller.NewYayasanController(yayasanService)
	router := httprouter.New()
	router.GET("/yayasans", yayasanController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/yayasans", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	var yayasans = responseBody["data"].([]interface{})
	defer db.Close()

	err1 := tmpl.Execute(w, yayasans)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetById(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "show.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}
	db := app.NewDB()
	validate := validator.New()
	yayasanRepository := repository.NewYayasanRepository()
	yayasanService := service.NewYayasanService(yayasanRepository, db, validate)
	yayasanController := controller.NewYayasanController(yayasanService)
	router := httprouter.New()
	router.GET("/show/:id", yayasanController.FindById)
	router.PanicHandler = exception.ErrorHandler
	fmt.Println("rRRRRRRBody", r.URL)
	Id := r.URL.Query().Get("id")
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/show/"+Id, nil)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, request)
	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//yayasan := domain.Yayasan{}
	yayasan := responseBody["data"]
	tmpl.Execute(w, yayasan)
}
