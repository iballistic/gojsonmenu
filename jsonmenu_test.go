package gojsonmenu

import (
	"encoding/json"
	"os"
	"testing"
)

func TestParse(t *testing.T) {

	jsonFile, err := os.ReadFile("test.json")
	if err != nil {
		t.Errorf("%s", err)
	}

	var menu JSonMenu
	err = json.Unmarshal(jsonFile, &menu)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %s", err)
	}

	config_data := menu.ViewGroupBySection("foodmenu")
	if len(config_data) != 3 {
		t.Errorf("Expected 3 sections, found %d", len(config_data))
	}
	for section, views := range config_data {
		t.Logf("Section: %s", section.Header)
		for _, view := range views {

			t.Logf("Cell Key: %s default value %s", view.Cell.Key, view.Cell.DefaultValue())

		}

	}

}

//https://go.dev/doc/code
//golang.org/x/example
