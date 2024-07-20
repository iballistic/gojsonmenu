package gojsonmenu

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

type ValueSelectionType string

const (
	// default value or optional values to provide a drop down or select menu
	// we want to display default value by default during new record creation
	// not all cells should have a value, but it is a good idea to have some values
	//for example windspeed could be 0 default
	DefaultValue  ValueSelectionType = "default"
	OptionalValue ValueSelectionType = "optional"
)

type ConversionSelectionType string

const (
	// DefaultUnit represents the default unit is used to store int or double/float values,
	//if a user does not select one of the optional supported conversions, the default value
	//should be stored as is, otherwise we must convert from optional
	// (normally we be a user setting stored somewhere else) unit back to a default unit
	//to ensure the data integrity. Otherwise, we must information in pairs, user selected unit plus value
	DefaultUnit  ConversionSelectionType = "default"
	OptionalUnit ConversionSelectionType = "optional"
)

type CellValue struct {
	Value string             `json:"value"`
	Type  ValueSelectionType `json:"type"`
}

type Converter struct {
	Name   string                  `json:"name"`
	Type   ConversionSelectionType `json:"type"`
	Symbol string                  `json:"symbol"`
}

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
	Converters    []Converter   `json:"converter"`
	Minval        string        `json:"minval"`
	Maxval        string        `json:"maxval"`
	Rowheight     int           `json:"rowheight"`
}

func (c *TableCell) DefaultValue() string {

	for _, item := range c.Values {
		if item.Type == DefaultValue {
			return item.Value
		}
	}

	return ""
}

func (c *TableCell) DefaultUnit() string {

	for _, item := range c.Converters {
		if item.Type == DefaultUnit {
			return item.Name
		}
	}

	return ""
}
