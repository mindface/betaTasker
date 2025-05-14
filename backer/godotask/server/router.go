package server

import (
	"github.com/godotask/controller/book"
	"github.com/godotask/controller/top"
	"github.com/godotask/controller/user"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(static.Serve("/usr/local/go/godotask/static", static.LocalFile("./images", true)))
	r.LoadHTMLGlob("view/*.html")
	r.Use(CORSMiddleware())

	r.POST("/api/login", user.Login) 

	r.Use(user.AuthMiddleware())

	r.GET("/", top.IndexDisplayAction)
	r.GET("/book", book.BookListDisplayAction)
	r.GET("/api/book", book.ApiBookListDisplayAction)
	r.GET("/book/add", book.BookAddDisplayAction)
	r.POST("/api/book", book.AddBookAction)
	r.POST("/api/file", book.HundleUplond)
	r.DELETE("/api/deletebook/:id", book.DeleteBookAction)
	r.PUT("/api/updatebook/:id", book.UpdateBookAction)
	r.GET("/book/edit/:id", book.UpdateBookAction)

	// User authentication routes
	r.POST("/api/register", user.Register)

	// Protected routes
	r.GET("/api/user/profile",  user.AuthMiddleware(), user.Profile)

	return r
}
