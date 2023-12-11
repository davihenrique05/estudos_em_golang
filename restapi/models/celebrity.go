package models

type Celebrity struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

var Celebrities []Celebrity
