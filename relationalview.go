package gojsonmenu

// Relational view based on mapping, which is defined in the Json file
type RelationalView struct {
	Storyboard Storyboard
	Section    TableSection
	Cell       TableCell
	Order      int
	Readonly   bool
}
