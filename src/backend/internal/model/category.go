package model

type Category struct {
	ID       int `json:"id"`
	ParentID int `json:"parent_id"`

	Title     string `json:"title"`
	MetaTitle string `json:"meta_title"`
	Slug      string `json:"slug"`
	Content   string `json:"content"`
}
