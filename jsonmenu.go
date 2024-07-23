package gojsonmenu

import (
	"errors"
	"sort"
)

type JSonMenu struct {
	Storyboard []Storyboard        `json:"storyboard"`
	Section    []TableSection      `json:"section"`
	Cells      []TableCell         `json:"cells"`
	Mapping    []RelationalMapping `json:"mapping"`
}

// Builds relational view based on the mapping confugration is the json file
func (j *JSonMenu) View() []RelationalView {

	storyboardCollection := make(map[string]Storyboard)
	sectionCollection := make(map[string]TableSection)
	cellCollection := make(map[string]TableCell)

	var data []RelationalView

	for _, storyboard := range j.Storyboard {

		storyboardCollection[storyboard.Name] = storyboard
	}

	for _, section := range j.Section {

		sectionCollection[section.Name] = section
	}

	for _, cell := range j.Cells {

		cellCollection[cell.Key] = cell
	}

	for _, mapping := range j.Mapping {
		var view RelationalView
		view.Storyboard = storyboardCollection[mapping.Storyboard]
		view.Section = sectionCollection[mapping.Section]
		view.Cell = cellCollection[mapping.Cell]
		view.Order = mapping.Order
		view.Readonly = mapping.Readonly
		data = append(data, view)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Order > data[j].Order
	})

	return data

}

// Parameter forStoryboard: Name of the story board that view will be filtered on
// Returns: An optional array of RelationalView for a story board(page)
func (j *JSonMenu) ViewByStoryboard(name string) []RelationalView {

	var data []RelationalView
	for _, view := range j.View() {
		if view.Storyboard.Name == name {
			data = append(data, view)
		}

	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Order < data[j].Order
	})

	return data
}

// Returns relational view for a story board (page)
// Parameter forStoryboard: Name of the story board that view will be filtered on
// Returns: An dictionary(map) of RelationalView, the dictionary(map) keys are TableSection object
// TableSection key is mostly for groups, a page(storyborard) may have multiple sections and each section may have
// different fields
func (j *JSonMenu) ViewGroupBySection(storyboard string) map[TableSection][]RelationalView {

	data := make(map[TableSection][]RelationalView)
	for _, view := range j.ViewByStoryboard(storyboard) {
		data[view.Section] = append(data[view.Section], view)
	}

	for key, val := range data {
		//sort items within val
		sort.Slice(val, func(i, j int) bool {
			return val[i].Order < val[j].Order
		})
		data[key] = val

	}
	return data
}

// Returns a single cell config by key
// Parameter byKey: byKey is the key name of a cell
// Returns: first elment in the array of cell config ( usualy there should be only one cell with the same key)
func (j *JSonMenu) GetCell(byKey string) (TableCell, error) {

	for _, cell := range j.Cells {
		if cell.Key == byKey {
			return cell, nil
		}

	}

	return TableCell{}, errors.New("not found")
}
