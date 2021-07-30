package models

type Article struct {
	Title     string `json:"title"`
	Intro     string `json:"intro"`
	Content   string `json:"content"`
	Author_id int    `json:"author_id"`
}
