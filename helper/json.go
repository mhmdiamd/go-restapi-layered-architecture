package helper

import (
	"encoding/json"
	"golang-resful-api/model/web"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, dataEntity interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(dataEntity)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, webResponse web.WebResponse) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
