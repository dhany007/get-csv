package delivery

import (
	"activity-csv/models/response"
	"encoding/csv"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (h handler) ProcessFruit(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// open csv
	file, err := os.Open("contoh.csv")
	if err != nil {
		response.Result(w, http.StatusBadRequest)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		response.Result(w, http.StatusBadRequest)
		return
	}

	result, err := h.fruitUsecase.ProcessCSV(data)
	if err != nil {
		response.Result(w, http.StatusBadRequest)
		return
	}

	response.ResultWithData(w, http.StatusOK, result)
}
