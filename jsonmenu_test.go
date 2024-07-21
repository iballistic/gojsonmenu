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

func TestTableCellValueUnit(t *testing.T) {
	tableCell := TableCell{
		Key:      "test",
		Title:    "My test",
		Format:   "%.3f",
		Celltype: CellDoubleType,
	}

	values := make([]CellValue, 0)
	values = append(values, CellValue{Value: "168", Requirement: DefaultValue})
	values = append(values, CellValue{Value: "165", Requirement: OptionalValue})
	values = append(values, CellValue{Value: "167", Requirement: OptionalValue})
	//lets add another default value to see if we can confuse the system.
	values = append(values, CellValue{Value: "165", Requirement: DefaultValue})
	//set values
	tableCell.Values = values

	valueUnits := make([]ValueUnit, 0)
	valueUnits = append(valueUnits, ValueUnit{Name: "grain", Option: DefaultUnit, Symbol: "gr"})
	valueUnits = append(valueUnits, ValueUnit{Name: "grams", Option: OptionalUnit, Symbol: "g"})

	tableCell.Units = valueUnits

	for _, item := range tableCell.Values {
		t.Logf("Original value %s", item.Value)
	}

	convertedValues := tableCell.ConvertedValues("grain", "grams")

	for _, item := range convertedValues {
		t.Logf("Converted value %s", item.Value)
	}

}

//https://go.dev/doc/code
//golang.org/x/example
