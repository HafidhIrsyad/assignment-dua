package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		return
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	jsonData, _ := json.Marshal(response)
	writer.Header().Add("Content-Type", "application/json")
	_, errWrite := writer.Write(jsonData)
	if errWrite != nil {
		return
	}
}
