package config

import (
	"encoding/json"
	"os"
)

func Parse(path string, config interface{}) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}
