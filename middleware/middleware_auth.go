package middleware

import (
	"golang-resful-api/helper"
	"golang-resful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-API-Key") == "RAHASIA" {
		// Ok
		middleware.Handler.ServeHTTP(writer, req)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZE",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}

}
