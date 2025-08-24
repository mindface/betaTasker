package server

import (
	"github.com/godotask/controller/book"
	"github.com/godotask/controller/memory"
	"github.com/godotask/controller/top"
	"github.com/godotask/controller/user"
	"github.com/godotask/controller/task"
	"github.com/godotask/controller/assessment"
	"github.com/godotask/repository"
	"github.com/godotask/service"
	"github.com/godotask/model"
	// "github.com/godotask/middleware" // 一時的にコメントアウト

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// シンプルなCORSミドルウェア（デバッグ用）
func CORSMiddlewareSimple() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	
	// CORSミドルウェアのみ適用（他のミドルウェアは一旦コメントアウト）
	// r.Use(middleware.LoggingMiddleware())
	// r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(CORSMiddlewareSimple())
	// r.Use(middleware.RequestValidationMiddleware())
	// r.Use(middleware.RateLimitMiddleware())
	
	r.Use(static.Serve("/usr/local/go/godotask/static", static.LocalFile("./images", true)))
	r.LoadHTMLGlob("view/*.html")


  bookRepo := &repository.BookRepositoryImpl{DB: model.DB}
	bookService := &service.BookService{Repo: bookRepo}
	bookController := book.BookController{Service: bookService}

	memoryRepo := &repository.MemoryRepositoryImpl{DB: model.DB}
	memoryContextRepo := &repository.MemoryContextRepositoryImpl{DB: model.DB}
	memoryService := &service.MemoryService{
		Repo: memoryRepo,
		ContextRepo: memoryContextRepo,
	}
	memoryController := memory.MemoryController{Service: memoryService}

	taskRepo := &repository.TaskRepositoryImpl{DB: model.DB}
	taskService := &service.TaskService{Repo: taskRepo}
	taskController := task.TaskController{Service: taskService}

  assessmentRepo := &repository.AssessmentRepositoryImpl{DB: model.DB}
	assessmentService := &service.AssessmentService{Repo: assessmentRepo}
	assessmentController := assessment.AssessmentController{Service: assessmentService}

	// 認証不要のエンドポイント
	r.POST("/api/login", user.Login)
	r.POST("/api/register", user.Register)
	r.POST("/api/logout", user.Logout)
	
	// 一時的に認証を無効化（デバッグ用）
	
	r.GET("/", top.IndexDisplayAction)
		
	// Book API (CRUD)
	r.GET("/api/book", bookController.ListBooks)
	r.POST("/api/file", book.HundleUplond)
	r.POST("/api/book", bookController.AddBook)
	r.DELETE("/api/deletebook/:id", bookController.DeleteBook)
	r.PUT("/api/updatebook/:id", bookController.EditBook)

	// Task API (CRUD)
	r.POST("/api/task", taskController.AddTask)
	r.GET("/api/task", taskController.ListTasks)
	r.GET("/api/task/:id", taskController.GetTask)
	r.PUT("/api/task/:id", taskController.EditTask)
	r.DELETE("/api/task/:id", taskController.DeleteTask)

	// User profile
	r.GET("/api/user/profile", user.AuthMiddleware(), user.Profile)

	// Memory API (CRUD)
	r.POST("/api/memory", memoryController.AddMemory)
	r.GET("/api/memory", memoryController.ListMemories)
	r.GET("/api/memory/:id", memoryController.GetMemory)
	r.GET("/api/memory/context/:code", memoryController.GetMemoryContextByCode)
	r.GET("/api/memory/aid/:code", memoryController.GetMemoryAidByCode)
	r.PUT("/api/memory/:id", memoryController.EditMemory)
	r.DELETE("/api/memory/:id", memoryController.DeleteMemory)

	// Assessment API (CRUD)
	r.POST("/api/assessment", assessmentController.AddAssessment)
	r.GET("/api/assessment", assessmentController.ListAssessments)
	r.POST("/api/assessmentsForTaskUser", assessmentController.ListAssessmentsForTaskUser)
	r.GET("/api/assessment/:id", assessmentController.GetAssessment)
	r.PUT("/api/assessment/:id", assessmentController.EditAssessment)
	r.DELETE("/api/assessment/:id", assessmentController.DeleteAssessment)	

	// 404ハンドラー（一時的にコメントアウト）
	// r.NoRoute(middleware.NotFoundMiddleware())
	
	return r
}
