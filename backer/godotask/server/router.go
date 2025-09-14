package server

import (
	"github.com/godotask/controller/book"
	"github.com/godotask/controller/memory"
	"github.com/godotask/controller/top"
	"github.com/godotask/controller/user"
	"github.com/godotask/controller/task"
	"github.com/godotask/controller/assessment"
	"github.com/godotask/controller/heuristics"
	"github.com/godotask/controller/process_optimization"
	"github.com/godotask/controller/qualitative_label"
	"github.com/godotask/controller/knowledge_pattern"
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

	heuristicsRepo := &repository.HeuristicsRepositoryImpl{DB: model.DB}
	heuristicsService := &service.HeuristicsService{Repo: heuristicsRepo}
	heuristicsController := heuristics.HeuristicsController{Service: heuristicsService}

	processOptimizationRepo := &repository.ProcessOptimizationRepositoryImpl{DB: model.DB}
	processOptimizationService := &service.ProcessOptimizationService{Repo: processOptimizationRepo}
	processOptimizationController := process_optimization.ProcessOptimizationController{Service: processOptimizationService}

	qualitativeLabelRepo := &repository.QualitativeLabelRepositoryImpl{DB: model.DB}
	qualitativeLabelService := &service.QualitativeLabelService{Repo: qualitativeLabelRepo}
	qualitativeLabelController := qualitative_label.QualitativeLabelController{Service: qualitativeLabelService}

	knowledgePatternRepo := &repository.KnowledgePatternRepositoryImpl{DB: model.DB}
	knowledgePatternService := &service.KnowledgePatternService{Repo: knowledgePatternRepo}
	knowledgePatternController := knowledge_pattern.KnowledgePatternController{Service: knowledgePatternService}

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

	// Heuristics API (ML Pipeline & Analytics)
	r.POST("/api/heuristics/analyze", heuristicsController.Analyze)
	r.GET("/api/heuristics/analyze/:id", heuristicsController.GetAnalysis)
	r.POST("/api/heuristics/track", heuristicsController.TrackBehavior)
	r.GET("/api/heuristics/track/:user_id", heuristicsController.GetTrackingData)
	r.GET("/api/heuristics/insights", heuristicsController.ListInsights)
	r.GET("/api/heuristics/insights/:id", heuristicsController.GetInsight)
	r.GET("/api/heuristics/patterns", heuristicsController.DetectPatterns)
	r.POST("/api/heuristics/patterns/train", heuristicsController.TrainModel)

	// Process Optimization API (CRUD)
	r.POST("/api/process_optimization", processOptimizationController.AddProcessOptimization)
	r.GET("/api/process_optimization", processOptimizationController.ListProcessOptimizations)
	r.GET("/api/process_optimization/:id", processOptimizationController.GetProcessOptimization)
	r.PUT("/api/process_optimization/:id", processOptimizationController.EditProcessOptimization)
	r.DELETE("/api/process_optimization/:id", processOptimizationController.DeleteProcessOptimization)

	// Qualitative Label API (CRUD)
	r.POST("/api/qualitative_label", qualitativeLabelController.AddQualitativeLabel)
	r.GET("/api/qualitative_label", qualitativeLabelController.ListQualitativeLabels)
	r.GET("/api/qualitative_label/:id", qualitativeLabelController.GetQualitativeLabel)
	r.PUT("/api/qualitative_label/:id", qualitativeLabelController.EditQualitativeLabel)
	r.DELETE("/api/qualitative_label/:id", qualitativeLabelController.DeleteQualitativeLabel)

	// Knowledge Pattern API (CRUD)
	r.POST("/api/knowledge_pattern", knowledgePatternController.AddKnowledgePattern)
	r.GET("/api/knowledge_pattern", knowledgePatternController.ListKnowledgePatterns)
	r.GET("/api/knowledge_pattern/:id", knowledgePatternController.GetKnowledgePattern)
	r.PUT("/api/knowledge_pattern/:id", knowledgePatternController.EditKnowledgePattern)
	r.DELETE("/api/knowledge_pattern/:id", knowledgePatternController.DeleteKnowledgePattern)


	// 404ハンドラー（一時的にコメントアウト）
	// r.NoRoute(middleware.NotFoundMiddleware())
	
	return r
}
