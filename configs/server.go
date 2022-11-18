package configs

import (
	"activity-csv/services/utils"
	"fmt"
	"log"
	"net/http"
)

func (c *Config) InitServer() error {
	envPort := utils.GetEnv("APP_PORT", "11010")

	port := fmt.Sprintf(":%s", envPort)
	log.Println("Server Listening on port", port)

	err := http.ListenAndServe(port, c.Router)
	if err != nil {
		return err
	}

	return nil
}
