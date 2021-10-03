package swagger

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func generateJsonReponse(w http.ResponseWriter, message string) ([]byte, error) {
	var response = Response{Message: message}
	return json.Marshal(response)
}

func writeSucess(w http.ResponseWriter) (writer http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return w
}

func writeFail(w http.ResponseWriter) (writer http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	return w
}
