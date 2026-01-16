package main

import (
	"log"
	"net/http"

	"github.com/godotask/server"
	"github.com/godotask/model"
	"github.com/joho/godotenv"
	// "giner/calculation"
	// "giner/controller"
)

// type UserInfo struct {
// 	UserId  int    `json:"user_id"`
// 	UserUi  string `json:"user_ui"`
// 	Contens struct {
// 		Id      int `json:"id"`
// 		Title   int `json:"Title"`
// 		Body    int `json:"body"`
// 		LabelId int `json:"label_id"`
// 	} `json:"contens"`
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	model.InitDB()
	// bytes, err := ioutil.ReadFile("data.json")
	// if err != nil {
	// 	 log.Fatal(err)
	// }
	// calculation := calculation.New()
	// r := gin.Default()
	// r.LoadHTMLGlob("templates/*.html")
	// r.GET("/", func(c *gin.Context) {
	// c.HTML(200,"index.html", gin.H{})
	// 	c.JSON(200, gin.H{
	// 		"msg": fmt.Println(string(bytes)),
	// 	})
	// })
	r := server.GetRouter()

	r.Run(":8080")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
