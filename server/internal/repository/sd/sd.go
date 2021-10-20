package sd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"uart-mp3-player/internal/models"
)

const JSON_DATA = "/../../../init/sd.json"

func GetAll() (models.SD, error) {
	var model models.SD
	var err error

	jsonFile, err := os.Open(getBasePath() + JSON_DATA)
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

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
