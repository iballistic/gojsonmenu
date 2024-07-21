package gojsonmenu

type UnitOptionType string

const (
	// DefaultUnit represents the default unit is used to store int or double/float values,
	//if a user does not select one of the optional supported conversions, the default value
	//should be stored as is, otherwise we must convert from optional
	// (normally we be a user setting stored somewhere else) unit back to the default unit
	//to ensure the data integrity. Otherwise, the information must be stored in pairs, user selected unit plus value
	DefaultUnit  UnitOptionType = "default"
	OptionalUnit UnitOptionType = "optional"
)

type ValueUnit struct {
	Name   string         `json:"name"`
	Option UnitOptionType `json:"type"`
	Symbol string         `json:"symbol"`
}
