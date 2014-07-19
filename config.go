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

	err = setDefaultValue(config)
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

func setDefaultValue(config interface{}) error {
	return setDefaultValueToStruct(reflect.TypeOf(config), reflect.ValueOf(config))
}
