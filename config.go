package config

import (
	"encoding/json"
	"os"
	"reflect"
)

func Parse(path string, config interface{}) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	setDefaultValue(config)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}

func setDefaultValue(config interface{}) {
	setDefaultValueToStruct(reflect.TypeOf(config), reflect.ValueOf(config))
}
