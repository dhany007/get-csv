package configs

import (
	"activity-csv/services/delivery"
	"activity-csv/services/usecase"
)

func (c *Config) InitService() error {
	fruitUsecase := usecase.NewFruitUsecase()

	c.Router = delivery.NewHandler(fruitUsecase)

	return nil
}
