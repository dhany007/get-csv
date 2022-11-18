package services

import "activity-csv/models"

type FruitUsecase interface {
	ProcessCSV(data [][]string) (result []models.ShopingCart, err error)
}
