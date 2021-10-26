package sd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"uart-mp3-player/internal/models"
)

func GetAll() (models.SD, error) {
	var model models.SD
	var err error

	jsonPath := os.Getenv("SD_DATA_PATH")
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return model, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return model, err
	}

	json.Unmarshal(byteValue, &model)

	return model, nil
}
