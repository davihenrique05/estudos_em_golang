package models

type Celebrity struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
