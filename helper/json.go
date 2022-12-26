package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(request)
	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, responses interface{}) {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(responses)
	if err != nil {
		panic(err)
	}
}
