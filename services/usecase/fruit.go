package usecase

import (
	"activity-csv/models"
	"activity-csv/services"
	"activity-csv/services/utils"
)

type fruitUsecase struct {
}

func NewFruitUsecase() services.FruitUsecase {
	return &fruitUsecase{}
}

// ProcessCSV implements services.FruitUsecase
func (f fruitUsecase) ProcessCSV(data [][]string) (result []models.ShopingCart, err error) {
	headers := []models.Object{}

	for i, line := range data {
		if i == 0 {
			// generate header
			headers = utils.GenerateHeader(line)
		} else {
			// fill data based on header
			fruit := utils.FillData(line, headers)
			result = append(result, fruit)
		}
	}

	return
}
