package models

type Book struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
