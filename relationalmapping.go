package gojsonmenu

// Mapping is defined in the Json file to build a relation between
type RelationalMapping struct {
	Storyboard string `json:"storyboard"`
	Section    string `json:"section"`
	Cell       string `json:"cell"`
	Order      int    `json:"order"`
	Readonly   bool   `json:"readonly"`
}
