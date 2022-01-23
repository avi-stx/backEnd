package config

import (
	"encoding/json"
	"io/ioutil"
)

type configProperties struct {
	TargetFolder string `json:"target_folder"`
}

var AppConfig = &configProperties{}

func ReadConfig(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, AppConfig)
	return err
}
