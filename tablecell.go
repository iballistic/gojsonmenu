package gojsonmenu

import (
	"fmt"
	"strconv"
)

type TableCellType string

const (
	//TableCellType is used in views, for example a from, there could different input fields
	CellStringType      TableCellType = "CellString"
	CellSegueType       TableCellType = "CellSegue" //is used to provide a link to jump to another page, to search and select a values
	CellDoubleType      TableCellType = "CellDouble"
	CellIntType         TableCellType = "CellInt"
	CellLargeStringType TableCellType = "CellLargeString"
	CellPhotoType       TableCellType = "CellPhoto"
)

type TableCell struct {
	Maxchar       int           `json:"maxchar"`
	Format        string        `json:"format"`
	Pattern       string        `json:"pattern"`
	Incrementstep string        `json:"incrementstep"`
	Key           string        `json:"key"`
	Title         string        `json:"title"`
	Celltype      TableCellType `json:"celltype"`
	Placeholder   string        `json:"placeholder"`
	Values        []CellValue   `json:"values"`
	Units         []ValueUnit   `json:"units"`
	Minval        string        `json:"minval"`
	Maxval        string        `json:"maxval"`
	Rowheight     int           `json:"rowheight"`
}

// / Returns: new title using current title and valut unit
func (c *TableCell) TitleComputed(unit *ValueUnit) string {
	if len(c.Title) == 0 {
		return ""
	}

	if unit != nil {
		return fmt.Sprintf("%s (%s)", c.Title, unit.Symbol)
	} else if c.DefaultUnit() != nil {
		return fmt.Sprintf("%s (%s)", c.Title, c.DefaultUnit().Symbol)

	} else {
		return c.Title
	}
}

// Computes new placeholder text using min and max values if applicable for provided unit
func (c *TableCell) PlaceholderComputed(unit *ValueUnit) string {

	if len(c.Placeholder) == 0 {
		return ""
	} else if unit == nil {
		return c.Placeholder
	} else if c.Celltype != CellIntType && c.Celltype != CellDoubleType {
		return c.Placeholder
	} else if len(c.Minval) == 0 || len(c.Maxval) == 0 {
		return c.Placeholder
	} else if c.Celltype == CellIntType {
		return fmt.Sprintf("%s (%s)", c.Placeholder, unit.Symbol)
	} else if c.Celltype == CellDoubleType {

		minVal, err := c.MinvalConverted(unit.Name)

		if err != nil {
			return c.Placeholder
		}

		maxVal, err := c.MaxvalConverted(unit.Name)

		if err != nil {
			return c.Placeholder
		}

		return fmt.Sprintf("%s (From %s To %s)", c.Placeholder, minVal, maxVal)

	} else {
		return c.Placeholder
	}
}

func (c *TableCell) DefaultValue() string {

	for _, item := range c.Values {
		if item.Requirement == DefaultValue {
			return item.Value
		}
	}

	return ""
}

func (c *TableCell) DefaultUnit() *ValueUnit {

	for _, item := range c.Units {
		if item.Option == DefaultUnit {
			return &item
		}
	}

	return nil
}

// Converts the default value per user setting
func (c *TableCell) DefaultValueConverted(to string) (string, error) {

	value := c.DefaultValue()
	defaultUnit := c.DefaultUnit()

	valueProper, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return value, err
	}

	convertedValue, err := ConvertValue(valueProper, defaultUnit.Name, to)
	if err != nil {
		return value, err
	}

	return fmt.Sprintf(c.Format, convertedValue), nil
}

// Converts the minimum value per user setting
func (c *TableCell) MinvalConverted(to string) (string, error) {

	value := c.Minval
	defaultUnit := c.DefaultUnit()

	valueProper, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return value, err
	}

	convertedValue, err := ConvertValue(valueProper, defaultUnit.Name, to)
	if err != nil {
		return value, err
	}

	return fmt.Sprintf(c.Format, convertedValue), nil

}

// Converts the maximum value per user setting
func (c *TableCell) MaxvalConverted(to string) (string, error) {
	value := c.Maxval
	defaultUnit := c.DefaultUnit()

	valueProper, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return value, err
	}

	convertedValue, err := ConvertValue(valueProper, defaultUnit.Name, to)
	if err != nil {
		return value, err
	}

	return fmt.Sprintf(c.Format, convertedValue), nil
}

// Converts all available values per user settings, in this case it would be toUnit
// mostly used for test cases
func (c *TableCell) ConvertedValues(fromUnit string, toUnit string) []CellValue {
	var convertedValues = make([]CellValue, 0)

	for _, item := range c.Values {

		newCellValue := CellValue{
			Requirement: item.Requirement,
			Value:       item.Value,
		}

		valueProper, err := strconv.ParseFloat(item.Value, 64)

		if err != nil {
			valueProper = 0.0
		}

		convertedValue, err := ConvertValue(valueProper, fromUnit, toUnit)
		if err != nil {
			newCellValue.Value = fmt.Sprintf(c.Format, 0)
		} else {
			newCellValue.Value = fmt.Sprintf(c.Format, convertedValue)
		}

		convertedValues = append(convertedValues, newCellValue)

	}

	return convertedValues
}
