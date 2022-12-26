package main

import (
	"belajarlagi/app"
	"belajarlagi/controller"
	"belajarlagi/middleware"
	"belajarlagi/repository"
	"belajarlagi/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	yayasanRepository := repository.NewYayasanRepository()
	yayasanService := service.NewYayasanService(yayasanRepository, db, validate)
	yayasanController := controller.NewYayasanController(yayasanService)
	router := app.NewRouter(yayasanController)

	server := http.Server{
		Addr:    "localhost:3030",
		Handler: middleware.NewAuthMiddleware(router),
	}
	server.ListenAndServe()
}
