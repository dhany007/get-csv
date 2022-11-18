package delivery

import (
	"activity-csv/models/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h handler) Ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response.Result(w, http.StatusOK)
}
