package main

import (
	"belajarlagi/httpHandler"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	//db := app.NewDB()
	//validate := validator.New()
	//yayasanRepository := repository.NewYayasanRepository()
	//yayasanService := service.NewYayasanService(yayasanRepository, db, validate)
	//yayasanController := controller.NewYayasanController(yayasanService)
	//router := app.NewRouter(yayasanController)

	server := http.Server{
		Addr: "localhost:9000",
		//Handler: middleware.NewAuthMiddleware(router),
	}
	http.HandleFunc("/", httpHandler.GetAll)
	http.HandleFunc("/show", httpHandler.GetById)
	server.ListenAndServe()
}
