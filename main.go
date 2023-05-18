package main

import (
	"golang-resful-api/app"
	"golang-resful-api/controller"
	"golang-resful-api/helper"
	"golang-resful-api/middleware"
	"golang-resful-api/repository"
	"golang-resful-api/service"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Initialize Router
	var router = app.NewRouter(categoryController)

	// Create server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	// Runing Server
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
