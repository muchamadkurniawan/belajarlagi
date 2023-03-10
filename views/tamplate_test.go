package views

import (
	"belajarlagi/app"
	"belajarlagi/controller"
	"belajarlagi/exception"
	"belajarlagi/repository"
	"belajarlagi/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"text/template"
)

func TestIndexHtml(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("../views", "index.html")
		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "learning golang",
			"name":  "Batman",
		}

		err1 := tmpl.Execute(w, data)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func TestLayoutViewHtml(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		db := app.NewDB()
		tx, _ := db.Begin()
		yayasanRepo := repository.NewYayasanRepository()
		all := yayasanRepo.GetAll(context.Background(), tx)
		fmt.Fprint(w, all)
		tmpl := template.Must(template.ParseFiles("layout.html"))
		tmpl.ExecuteTemplate(w, "Index", all)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
func TestLayoutViewLayout(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db := app.NewDB()
		//tx, _ := db.Begin()
		validate := validator.New()
		yayasanRepo := repository.NewYayasanRepository()
		yayasanService := service.NewYayasanService(yayasanRepo, db, validate)
		yayasanController := controller.NewYayasanController(yayasanService)

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
		var yayasans = responseBody["data"].([]interface{})

		//fmt.Fprint(w, yayasans)
		tmpl := template.Must(template.ParseFiles("layout.html"))
		tmpl.ExecuteTemplate(w, "Index", yayasans)

	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func TestViewIndex2(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("../views", "index2.html")
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
		router.GET("/api/yayasans", yayasanController.FindAll)

		router.PanicHandler = exception.ErrorHandler

		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/yayasans", nil)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)
		var yayasans = responseBody["data"].([]interface{})

		err1 := tmpl.Execute(w, yayasans)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/show", Show)
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func TestGetID_show(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("../views", "show.html")
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
		router.GET("/api/yayasan/:id", yayasanController.FindById)

		router.PanicHandler = exception.ErrorHandler

		request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/api/yayasan/2", nil)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)
		//yayasan := domain.Yayasan{}
		yayasan := responseBody["data"]
		tmpl.Execute(w, yayasan)
	})
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func Show(w http.ResponseWriter, r *http.Request) {
	//db := app.NewDB()
	nId := r.URL.Query().Get("id")
	fmt.Println(nId)
	fmt.Fprint(w, nId)
}

func TestYayasanControllerImpl_FindAll(t *testing.T) {
	db := app.NewDB()

	validate := validator.New()
	yayasanRepository := repository.NewYayasanRepository()
	yayasanService := service.NewYayasanService(yayasanRepository, db, validate)
	yayasanController := controller.NewYayasanController(yayasanService)
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

func TestYayasanControllerImpl_FindByID(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepository := repository.NewYayasanRepository()
	yayasanService := service.NewYayasanService(yayasanRepository, db, validate)
	yayasanController := controller.NewYayasanController(yayasanService)
	router := httprouter.New()
	router.GET("/api/yayasan/:id", yayasanController.FindById)

	router.PanicHandler = exception.ErrorHandler

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9000/api/yayasan/1", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	//yayasan := domain.Yayasan{}
	yayasan := responseBody["data"]
	fmt.Println(yayasan)
}
