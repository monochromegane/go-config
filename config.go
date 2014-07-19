package config

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"
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

	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)
	for i := 0; i < t.Elem().NumField(); i++ {
		f := t.Elem().FieldByIndex([]int{i})
		def := f.Tag.Get("default")
		value := v.Elem().FieldByName(f.Name)
		switch f.Type.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, _ := strconv.ParseInt(def, 0, f.Type.Bits())
			value.SetInt(iv)
		case reflect.String:
			value.SetString(def)
		}
	}
}
