package config

import "testing"

type ExampleConfig struct {
	Name string `json:"name"`
}

func TestParse(t *testing.T) {

	config := &ExampleConfig{}
	err := Parse("examples/example.json", config)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if config.Name != "go-config" {
		t.Errorf("It should be equal %s, but %s", "go-config", config.Name)
	}
}

func TestSetDefalutValue(t *testing.T) {

	var conf = struct {
		String string `default:"default"`
		Int    int    `default:"10"`
		Int8   int8   `default:"20"`
		Int16  int16  `default:"30"`
		Int32  int32  `default:"40"`
		Int64  int64  `default:"50"`
	}{}

	setDefaultValue(&conf)

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

}
