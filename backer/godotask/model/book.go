package model

import (
	// "database/sql"
	// "encoding/json"
	// "fmt"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
)

//  create table book(id int not null primary key auto_increment, title text, name varchar(299), text text,disc text,imgPath varchar(12999), status varchar(299));
//  INSERT INTO book (title, name, text, disc, imgPath, status) values("title01","name","text","disc","imgPath","run");

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

// func GetBookList() string {

// 	db, err := sql.Open("mysql", bookname+":"+password+"@tcp(dbgodotask:3306)/"+schema)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("select id, title, name, text, disc, imgPath from book order by id")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer rows.Close()

// 	// list := make(map[int]string)
// 	var uptData []BookApi

// 	for rows.Next() {
// 		var book BookApi

// 		err := rows.Scan(&book.Id, &book.Title, &book.Name, &book.Text, &book.Disc, &book.ImgPath)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// list[book.id] = book.name
// 		fmt.Println(book.Id)
// 		inuptData := BookApi{Id: book.Id, Title: book.Title, Name: book.Name, Text: book.Text, Disc: book.Disc, ImgPath: book.ImgPath}
// 		uptData = append(uptData, inuptData)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(uptData)
// 	bytes, _ := json.Marshal(uptData)
// 	// if bytes != nil {
// 	// 	return "[]"
// 	// }

// 	return string(bytes)
// }

// func GetBookData(id string) map[string]string {
// 	db, err := sql.Open("mysql", bookname+":"+password+"@tcp(dbgodotask:3306)/"+schema)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("select id, name from user where id = ?", id)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	data := make(map[string]string)

// 	for rows.Next() {
// 		var book Book

// 		err := rows.Scan(&book.id, &book.name)
// 		if err != nil {
// 			panic(err)
// 		}

// 		data["id"] = strconv.Itoa(book.id)
// 		data["name"] = book.name
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return data
// }

// func EditBookData(id string, title string, name string, text string, disc string, imgPath string) {
// 	db, err := sql.Open("mysql", bookname+":"+password+"@tcp(dbgodotask:3306)/"+schema)
// 	fmt.Printf(id)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	update, err := db.Prepare("update book set title = ?, text = ?, name = ?, disc = ?, imgPath = ? where id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer update.Close()

// 	_, err = update.Exec(title, text, name, disc, imgPath, id)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// func DeleteBookData(id string) {
// 	db, err := sql.Open("mysql", bookname+":"+password+"@tcp(dbgodotask:3306)/"+schema)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	delete, err := db.Prepare("delete from book where id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer delete.Close()

// 	_, err = delete.Exec(id)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// func AddBookData(id int, title string, name string, text string, disc string, imgPath string) {
// 	// snedData := Book{}
// 	// snedData.name = name
// 	// snedData.title = title
// 	// snedData.text = text
// 	// snedData.disc = disc
// 	// snedData.imgPath = imgPath
// 	db, err := sql.Open("mysql", bookname+":"+password+"@tcp(dbgodotask:3306)/"+schema)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	insert, err := db.Prepare(`insert book(title,name,text,disc,imgPath) values(?,?,?,?,?)`)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()

// 	_, err = insert.Exec(title, name, text, disc, imgPath)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
