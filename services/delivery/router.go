package delivery

import (
	"activity-csv/services"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	fruitUsecase services.FruitUsecase
}

func NewHandler(
	fruitUsecase services.FruitUsecase,
) (router *httprouter.Router) {
	router = httprouter.New()

	h := handler{
		fruitUsecase,
	}

	// ping router
	router.GET("/ping", h.Ping)

	// dummy csv
	router.POST("/dummy/csv", h.ProcessFruit)

	return
}
