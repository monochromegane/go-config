package config

import "testing"

type ExampleConfig struct {
	Name string `json:name`
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
