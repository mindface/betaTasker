package model

const bookname = "root"
const password = "dbgodotask"
const schema = "dbgodotask"

type Book struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Name    string `json:"name"`
	Text    string `json:"text"`
	Disc    string `json:"disc"`
	ImgPath string `json:"imgPath"`
	Status  string `json:"status"`
}

func (Book) TableName() string {
	return "book"
}
