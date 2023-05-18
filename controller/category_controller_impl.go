package controller

import (
	"encoding/json"
	"golang-resful-api/helper"
	"golang-resful-api/model/web"
	"golang-resful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// decode data stream dari request body
	decoder := json.NewDecoder(req.Body)
	// Initialize variable categoryCreateRequest, this variable will send to service
	categoryCreateRequest := web.CategoryCreateRequest{}
	// Melakukan decoded pada variable yang berisi json data requset body dan memasukan isinya kedalam variable categoryCreateRequest
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	// Kirim data ke service
	response, err := controller.CategoryService.Create(req.Context(), categoryCreateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	writer.Header().Add("Content-Type", "application/json")
	// Membuat encoder dari response/writer
	encoder := json.NewEncoder(writer)
	// Ubah Data response ke bentuk json kembali
	err = encoder.Encode(&webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// get id From params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(req, &categoryUpdateRequest)
	categoryUpdateRequest.Id = id

	// Send to service
	response, err := controller.CategoryService.Update(req.Context(), categoryUpdateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	// canvert and send response json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// get id From params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(req.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	response := controller.CategoryService.FindById(req.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	response := controller.CategoryService.FindAll(req.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
