package model

import "time"

//  create table book(id int not null primary key auto_increment, title text, name varchar(299), text text,disc text,imgPath varchar(12999), status varchar(299));
//  INSERT INTO book (title, name, text, disc, imgPath, status) values("title01","name","text","disc","imgPath","run");

const bookname = "root"
const password = "dbgodotask"
const schema = "dbgodotask"

type Book struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	TaskID  int    `json:"task_id"`
	Title   string `json:"title"`
	Name    string `json:"name"`
	// 本を読んだ自分に関するメモ
	Text    string `json:"text"`
	// 本の要約や本の確認
	Disc    string `json:"disc"`
	// 本の画像パス(画像保存機能次第)
	ImgPath string `json:"imgPath"`
	// 本の状態(例: 読んでいる[run], 読み終わった[end], 読みたい[do], 読みたい[])
	Status  string `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (Book) TableName() string {
	return "book"
}
