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
	}{}

	setDefaultValue(&conf)

	if conf.String != "default" {
		t.Errorf("Expected String value 'defalut', but %s.", conf.String)
	}
	if conf.Int != 10 {
		t.Errorf("Expected String value 10, but %d.", conf.Int)
	}

}
