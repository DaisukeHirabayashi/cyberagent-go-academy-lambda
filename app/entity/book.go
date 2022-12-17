package entity

type Book struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
