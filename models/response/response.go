package response

import (
	"encoding/json"
	"net/http"
)

type ReturningValue struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func write(w http.ResponseWriter, result ReturningValue) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(result.Status)

	err := json.NewEncoder(w).Encode(&result)
	if err != nil {
		panic(err)
	}
}

func Result(w http.ResponseWriter, code int) {
	ResultWithData(w, code, nil)
}

func ResultWithData(w http.ResponseWriter, code int, data interface{}) {
	result := ReturningValue{
		Status: code,
		Data:   data,
	}

	write(w, result)
}
