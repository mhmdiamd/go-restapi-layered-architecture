package test

import (
	"golang-resful-api/app"
	"golang-resful-api/controller"
	"golang-resful-api/middleware"
	"golang-resful-api/repository"
	"golang-resful-api/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupRouter() http.Handler {
	db := SetupTestDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Initialize Router
	var router = app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()

	// Create request Body
	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get Result
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T) {
	router := setupRouter()

	// Create request Body
	requestBody := strings.NewReader(`{"name" : "Fashion"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/5", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get Result
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestFindAllCategorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestFindByIdCategorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCategorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestUnauthorizeSuccess(t *testing.T) {
	router := setupRouter()

	// Create request Body
	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get Result
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestCreateCategoryValidationSuccess(t *testing.T) {
	router := setupRouter()

	// Create request Body
	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get Result
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestFindByIdCategoryNotFound(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/9", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}
