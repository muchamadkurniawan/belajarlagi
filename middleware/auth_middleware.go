package middleware

import (
	"belajarlagi/helper"
	"belajarlagi/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Handler.ServeHTTP(writer, request)
	webResponse := web.WebResponse{}
	helper.WriteToResponseBody(writer, webResponse)
}
