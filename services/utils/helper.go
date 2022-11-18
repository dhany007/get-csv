package utils

import (
	"activity-csv/models"
	"activity-csv/models/constant"
	"encoding/json"
	"os"
)

func GetEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func FillData(line []string, headers []models.Object) (result models.ShopingCart) {
	tmps := map[string]interface{}{}
	for i, field := range line {
		for _, head := range headers {
			if i == head.Key {
				tmps[head.Value] = field
			}
		}
	}

	jsonStr, err := json.Marshal(tmps)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonStr, &result)
	if err != nil {
		return
	}

	return
}

func GenerateHeader(line []string) (result []models.Object) {
	obj := models.Object{}

	for i, field := range line {
		for _, head := range constant.HEADER {
			if field == head {
				obj.Key = i
				obj.Value = field

				result = append(result, obj)
			}
		}
	}

	return
}
