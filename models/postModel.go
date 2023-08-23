package models

type Post struct {
	CustomModel
	UserID uint   `gorm:"user_id" json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Likes  int    `json:"likes"`
	Draft  bool   `json:"draft"`
	Author string `json:"author"`
}
