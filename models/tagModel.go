package models

type Tag struct {
	CustomModel
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) GetName() string {
	return "blog_tag"
}
