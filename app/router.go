package app

import (
	"belajarlagi/controller"
	"belajarlagi/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(yayasanController controller.YayasanController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/yayasans", yayasanController.FindAll)
	router.GET("/api/categories/:categoryId", yayasanController.FindById)
	router.POST("/api/categories", yayasanController.Create)
	router.PUT("/api/categories/:categoryId", yayasanController.Update)
	router.DELETE("/api/categories/:categoryId", yayasanController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
