package models

type Order struct {
	Items int `json:"items"`
}

type Pack struct {
	Size  int `json:"size"`
	Count int `json:"count"`
}
