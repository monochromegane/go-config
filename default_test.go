package config

import "testing"

func TestDefaultValueToBasicType(t *testing.T) {

	var conf = struct {
		Bool    bool    `default:"true"`
		String  string  `default:"default"`
		Int     int     `default:"10"`
		Int8    int8    `default:"20"`
		Int16   int16   `default:"30"`
		Int32   int32   `default:"40"`
		Int64   int64   `default:"50"`
		Uint    uint    `default:"10"`
		Uint8   uint8   `default:"20"`
		Uint16  uint16  `default:"30"`
		Uint32  uint32  `default:"40"`
		Uint64  uint64  `default:"50"`
		Float32 float32 `default:"10.1"`
		Float64 float64 `default:"20.2"`
	}{}

	setDefaultValue(&conf)

	if !conf.Bool {
		t.Errorf("Expected String value to be true")
	}
	if conf.String != "default" {
		t.Errorf("Expected String value 'defalut', but %s.", conf.String)
	}
	if conf.Int != 10 {
		t.Errorf("Expected Int value 10, but %d.", conf.Int)
	}
	if conf.Int8 != 20 {
		t.Errorf("Expected Int8 value 20, but %d.", conf.Int8)
	}
	if conf.Int16 != 30 {
		t.Errorf("Expected Int16 value 30, but %d.", conf.Int16)
	}
	if conf.Int32 != 40 {
		t.Errorf("Expected Int32 value 40, but %d.", conf.Int32)
	}
	if conf.Int64 != 50 {
		t.Errorf("Expected Int64 value 50, but %d.", conf.Int64)
	}
	if conf.Uint != 10 {
		t.Errorf("Expected Uint value 10, but %d.", conf.Uint)
	}
	if conf.Uint8 != 20 {
		t.Errorf("Expected Uint8 value 20, but %d.", conf.Uint8)
	}
	if conf.Uint16 != 30 {
		t.Errorf("Expected Uint16 value 30, but %d.", conf.Uint16)
	}
	if conf.Uint32 != 40 {
		t.Errorf("Expected Uint32 value 40, but %d.", conf.Uint32)
	}
	if conf.Uint64 != 50 {
		t.Errorf("Expected Uint64 value 50, but %d.", conf.Uint64)
	}
	if conf.Float32 != 10.1 {
		t.Errorf("Expected Float32 value 10.1, but %d.", conf.Float32)
	}
	if conf.Float64 != 20.2 {
		t.Errorf("Expected Float64 value 20.2, but %d.", conf.Float64)
	}
}

func TestNotDefineDefaultValue(t *testing.T) {
	var conf = struct {
		NoDefault int
	}{}
	setDefaultValue(&conf)

	if conf.NoDefault != 0 {
		t.Errorf("Expected NoDefault value 0, but %d.", conf.NoDefault)
	}
}

func TestDefaultValueToStruct(t *testing.T) {
	var conf = struct {
		Sub struct {
			Default string `default:"default"`
		}
	}{}

	setDefaultValue(&conf)
	if conf.Sub.Default != "default" {
		t.Errorf("Expected Sub.Default value default, but %s", conf.Sub.Default)
	}
}
