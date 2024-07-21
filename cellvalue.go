package gojsonmenu

type ValueRequirement string

const (
	// default value or optional values to provide a drop down or select menu
	// we want to display default value by default during new record creation
	// not all cells should have a value or default, but it is a good idea to have some values
	//for example windspeed could be 0 default
	DefaultValue  ValueRequirement = "default"
	OptionalValue ValueRequirement = "optional"
)

type CellValue struct {
	Value       string           `json:"value"`
	Requirement ValueRequirement `json:"type"`
}
