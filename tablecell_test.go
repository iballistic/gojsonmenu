package gojsonmenu

import (
	"testing"
)

func TestTableCellValueUnit(t *testing.T) {
	tableCell := TableCell{
		Key:      "test",
		Title:    "My test",
		Format:   "%.3f",
		Celltype: CellDoubleType,
	}

	values := make([]CellValue, 0)
	values = append(values, CellValue{Value: "10", Requirement: DefaultValue})
	values = append(values, CellValue{Value: "11", Requirement: OptionalValue})
	values = append(values, CellValue{Value: "12", Requirement: OptionalValue})
	//lets add another default value to see if we can confuse the system.
	values = append(values, CellValue{Value: "13", Requirement: DefaultValue})
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

func TestTableCellTitleComputed(t *testing.T) {
	tableCell := TableCell{
		Key:      "goldweight",
		Title:    "Gold Weight",
		Format:   "%.3f",
		Celltype: CellDoubleType,
	}

	values := make([]CellValue, 0)
	values = append(values, CellValue{Value: "10", Requirement: DefaultValue})
	values = append(values, CellValue{Value: "15", Requirement: OptionalValue})
	values = append(values, CellValue{Value: "16", Requirement: OptionalValue})
	//lets add another default value to see if we can confuse the system.
	values = append(values, CellValue{Value: "17", Requirement: DefaultValue})
	//set values
	tableCell.Values = values

	valueUnits := make([]ValueUnit, 0)
	valueUnits = append(valueUnits, ValueUnit{Name: "cararts", Option: DefaultUnit, Symbol: "ct"})
	valueUnits = append(valueUnits, ValueUnit{Name: "grams", Option: OptionalUnit, Symbol: "g"})
	valueUnits = append(valueUnits, ValueUnit{Name: "milligrams", Option: OptionalUnit, Symbol: "mg"})

	tableCell.Units = valueUnits

	newTitle := "Gold Weight (g)"
	titleComputed := tableCell.TitleComputed(&tableCell.Units[1])
	if newTitle != titleComputed {
		t.Errorf("was expecting \"%s\" got %s", newTitle, titleComputed)
	}
}

func TestTableCellPlaceHolderComputed(t *testing.T) {
	tableCell := TableCell{
		Key:         "windspeed",
		Title:       "Wind Speed",
		Placeholder: "Wind Speed",
		Format:      "%.2f",
		Minval:      "0",
		Maxval:      "70",
		Celltype:    CellDoubleType,
	}

	values := make([]CellValue, 0)
	values = append(values, CellValue{Value: "5", Requirement: DefaultValue})
	//set values
	tableCell.Values = values

	valueUnits := make([]ValueUnit, 0)
	valueUnits = append(valueUnits, ValueUnit{Name: "milesPerHour", Option: DefaultUnit, Symbol: "mph"})
	valueUnits = append(valueUnits, ValueUnit{Name: "kilometersPerHour", Option: OptionalUnit, Symbol: "km/h"})
	valueUnits = append(valueUnits, ValueUnit{Name: "footPerSecond", Option: OptionalUnit, Symbol: "ft/s"})

	tableCell.Units = valueUnits

	newPlaceholder := "Wind Speed (From 0.00 To 102.67)"
	placeholderComputed := tableCell.PlaceholderComputed(&tableCell.Units[2])
	if placeholderComputed != newPlaceholder {
		t.Errorf("was expecting \"%s\" got %s", newPlaceholder, placeholderComputed)
	}
}
