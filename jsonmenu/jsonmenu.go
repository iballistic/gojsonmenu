package jsonmenu

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

type Storyboard struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type TableSection struct {
	Name   string `json:"name"`
	Header string `json:"header"`
	Footer string `json:"footer"`
	Order  int    `json:"order"`
}

type CellValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Converter struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Symbol string `json:"symbol"`
}

type TableCell struct {
	Maxchar     int         `json:"maxchar"`
	Format      string      `json:"format"`
	Key         string      `json:"key"`
	Title       string      `json:"title"`
	Celltype    string      `json:"celltype"`
	Placeholder string      `json:"placeholder"`
	Values      []CellValue `json:"values"`
	Converters  []Converter `json:"converter"`
	Minval      string      `json:"minval"`
	Maxval      string      `json:"maxval"`
	Rowheight   int         `json:"rowheight"`
}

type RelationalMapping struct {
	Storyboard string `json:"storyboard"`
	Section    string `json:"section"`
	Cell       string `json:"cell"`
	Order      int    `json:"order"`
	Readonly   bool   `json:"readonly"`
}

type RelationalView struct {
	Storyboard Storyboard
	Section    TableSection
	Cell       TableCell
	Order      int
	Readonly   bool
}

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

func (j *JSonMenu) ViewByStoryboard(name string) []RelationalView {

	var data []RelationalView
	for _, view := range j.View() {
		if view.Storyboard.Name == name {
			data = append(data, view)
		}

	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Order > data[j].Order
	})

	return data
}

func (j *JSonMenu) ViewGroupBySection(storyboard string) map[TableSection][]RelationalView {

	data := make(map[TableSection][]RelationalView)
	for _, view := range j.ViewByStoryboard(storyboard) {
		data[view.Section] = append(data[view.Section], view)
	}

	// sort.Slice(data, func(i, j int) bool {
	// 	return data[i].Order > data[j].Order
	// })

	return data
}

func (j *JSonMenu) ViewByCell(key string) (TableCell, error) {

	for _, cell := range j.Cells {
		if cell.Key == key {
			return cell, nil
		}

	}

	return TableCell{}, errors.New("not found")
}
