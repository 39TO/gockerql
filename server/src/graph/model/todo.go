package model

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}
