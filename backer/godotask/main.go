package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/godotask/server"
	"github.com/godotask/model"
	"github.com/godotask/seed"
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

type UserInfo struct {
	UserId  int    `json:"user_id"`
	UserUi  string `json:"user_ui"`
	Contens string `json:"contens"`
}

type jsonData struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

var dataInfo = []UserInfo{{
	UserId:  0,
	UserUi:  "standard",
	Contens: "testtesttesttesttesttesttesttesttesttesttest",
}, {
	UserId:  1,
	UserUi:  "standard",
	Contens: "testtesttesttesttesttesttesttesttesttesttest",
}}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	model.InitDB()

	// Check if seed flag is provided
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		log.Println("Running database seeding...")
		if err := seed.RunAllSeeds(); err != nil {
			log.Fatalf("Seeding failed: %v", err)
		}
		log.Println("Seeding completed successfully!")
		return
	}

	// Check if clean-seed flag is provided
	if len(os.Args) > 1 && os.Args[1] == "clean-seed" {
		log.Println("Cleaning database and running seeding...")
		if err := seed.CleanAndSeed(); err != nil {
			log.Fatalf("Clean and seed failed: %v", err)
		}
		log.Println("Clean and seed completed successfully!")
		return
	}
	// bytes, err := ioutil.ReadFile("data.json")
	// if err != nil {
	// 	 log.Fatal(err)
	// }
	// calculation := calculation.New()
	hander001 := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&dataInfo); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}
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

	http.HandleFunc("/json", hander001)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
