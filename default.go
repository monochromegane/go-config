package config

import (
	"reflect"
	"strconv"
)

func setDefaultValueToStruct(t reflect.Type, v reflect.Value) error {

	var te reflect.Type
	if isPtrTypeOfStruct(t) {
		te = t.Elem()
	} else {
		te = t
	}

	var ve reflect.Value
	if isPtrValueOfStruct(v) {
		ve = v.Elem()
	} else {
		ve = v
	}

	for i := 0; i < te.NumField(); i++ {
		f := te.FieldByIndex([]int{i})
		def := f.Tag.Get("default")
		value := ve.FieldByName(f.Name)

		if f.Type.Kind() == reflect.Struct {
			return setDefaultValueToStruct(f.Type, value)
		}

		if def == "" {
			continue
		}

		switch f.Type.Kind() {
		case reflect.String:
			value.SetString(def)
		case reflect.Bool:
			bv, err := strconv.ParseBool(def)
			if err != nil {
				return err
			}
			value.SetBool(bv)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(def, 0, f.Type.Bits())
			if err != nil {
				return err
			}
			value.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uv, err := strconv.ParseUint(def, 0, f.Type.Bits())
			if err != nil {
				return err
			}
			value.SetUint(uv)
		case reflect.Float32, reflect.Float64:
			fv, err := strconv.ParseFloat(def, f.Type.Bits())
			if err != nil {
				return err
			}
			value.SetFloat(fv)
		}
	}
	return nil
}

func isPtrTypeOfStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func isPtrValueOfStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct
}
