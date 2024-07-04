package jsonmenu

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
		t.Errorf("%s", err)
	}

	config_data := menu.ViewGroupBySection("foodmenu")
	if len(config_data) != 3 {
		t.Errorf("%s", "found more or less than 3 sections")
	}
	for section, views := range config_data {
		t.Logf("Section: %s", section.Header)
		for _, view := range views {

			t.Logf("Cell Key: %s", view.Cell.Key)

		}

	}

}

//https://go.dev/doc/code
//golang.org/x/example
