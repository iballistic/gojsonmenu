package gojsonmenu

type TableSection struct {
	Name   string `json:"name"`
	Header string `json:"header"`
	Footer string `json:"footer"`
	Order  int    `json:"order"`
}
