package models

type Article struct {
	*CustomModel
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"coverImageUrl"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
