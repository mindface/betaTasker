package server

import (
	"github.com/godotask/controller/book"
	"github.com/godotask/controller/memory"
	"github.com/godotask/controller/user"
	"github.com/godotask/controller/task"
	"github.com/godotask/controller/assessment"
	"github.com/godotask/controller/heuristics"
	"github.com/godotask/controller/heuristics/analyze"
	"github.com/godotask/controller/heuristics/insight"
	"github.com/godotask/controller/process_optimization"
	"github.com/godotask/controller/qualitative_label"
	"github.com/godotask/controller/knowledge_pattern"
	"github.com/godotask/controller/language_optimization"
	"github.com/godotask/controller/teaching_free_control"
	"github.com/godotask/controller/phenomenological_framework"
	"github.com/godotask/service"
	"github.com/godotask/repository"
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
	// r.Use(CORSMiddlewareSimple())
	// r.Use(middleware.RequestValidationMiddleware())
	// r.Use(middleware.RateLimitMiddleware())
	
	r.Use(static.Serve("/usr/local/go/godotask/static", static.LocalFile("./images", true)))
	// r.LoadHTMLGlob("view/*.html")

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

	AnalyzeController := analyze.AnalyzeController{Service: heuristicsService}

	heuristicsInsightRepo := &repository.HeuristicsInsightRepositoryImpl{DB: model.DB}
	heuristicsInsightService := &service.HeuristicsInsightService{Repo: heuristicsInsightRepo}
	heuristicsInsightController := insight.HeuristicsInsightController{Service: heuristicsInsightService}



	processOptimizationRepo := &repository.ProcessOptimizationRepositoryImpl{DB: model.DB}
	processOptimizationService := &service.ProcessOptimizationService{Repo: processOptimizationRepo}
	processOptimizationController := process_optimization.ProcessOptimizationController{Service: processOptimizationService}

	qualitativeLabelRepo := &repository.QualitativeLabelRepositoryImpl{DB: model.DB}
	qualitativeLabelService := &service.QualitativeLabelService{Repo: qualitativeLabelRepo}
	qualitativeLabelController := qualitative_label.QualitativeLabelController{Service: qualitativeLabelService}

	knowledgePatternRepo := &repository.KnowledgePatternRepositoryImpl{DB: model.DB}
	knowledgePatternService := &service.KnowledgePatternService{Repo: knowledgePatternRepo}
	knowledgePatternController := knowledge_pattern.KnowledgePatternController{Service: knowledgePatternService}

	LanguageOptimizationRepo := &repository.LanguageOptimizationRepositoryImpl{DB: model.DB}
	LanguageOptimizationService := &service.LanguageOptimizationService{Repo: LanguageOptimizationRepo}
	LanguageOptimizationController := language_optimization.LanguageOptimizationController{Service: LanguageOptimizationService}

	TeachingFreeControlRepo := &repository.TeachingFreeControlRepositoryImpl{DB: model.DB}
	TeachingFreeControlService := &service.TeachingFreeControlService{Repo: TeachingFreeControlRepo}
	TeachingFreeControlController := teaching_free_control.TeachingFreeControlController{Service: TeachingFreeControlService}

	phenomenologicalFrameworkRepo := &repository.PhenomenologicalFrameworkRepositoryImpl{DB: model.DB}
	phenomenologicalFrameworkService := &service.PhenomenologicalFrameworkService{Repo: phenomenologicalFrameworkRepo}
	phenomenologicalFrameworkController := phenomenological_framework.PhenomenologicalFrameworkController{Service: phenomenologicalFrameworkService}

	// 認証不要のエンドポイント
	r.POST("/api/login", user.Login)
	r.POST("/api/register", user.Register)
	r.POST("/api/logout", user.Logout)
	
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
	r.POST("/api/heuristics/analyze", AnalyzeController.AddAnalyzeData)
	r.GET("/api/heuristics/analyze/:id", AnalyzeController.GetAnalyzeData)
	r.PUT("/api/heuristics/analyze/:id", AnalyzeController.EditAnalyzeData)
	r.DELETE("/api/heuristics/analyze/:id", AnalyzeController.DeleteAnalyzeData)

	// Heuristics API (ML Pipeline & Analytics)
	r.POST("/api/heuristics/insight", heuristicsInsightController.AddInsightData)
	r.GET("/api/heuristics/insight/:id", heuristicsInsightController.GetInsightData)
	r.PUT("/api/heuristics/insight/:id", heuristicsInsightController.EditInsightsData)
	r.DELETE("/api/heuristics/insight/:id", heuristicsInsightController.DeleteInsightData)

	r.POST("/api/heuristics/track", heuristicsController.TrackBehavior)
	r.GET("/api/heuristics/track/:user_id", heuristicsController.GetTrackingData)
	r.GET("/api/heuristics/patterns", heuristicsController.DetectPatterns)
	r.POST("/api/heuristics/patterns/train", heuristicsController.TrainModel)

	// phenomenological framework API (CRUD)
	r.POST("/api/phenomenological_framework", phenomenologicalFrameworkController.AddPhenomenologicalFramework)
	r.GET("/api/phenomenological_framework", phenomenologicalFrameworkController.ListPhenomenologicalFrameworks)
	r.GET("/api/phenomenological_framework/:id", phenomenologicalFrameworkController.GetPhenomenologicalFramework)
	r.PUT("/api/phenomenological_framework/:id", phenomenologicalFrameworkController.EditPhenomenologicalFramework)
	r.DELETE("/api/phenomenological_framework/:id", phenomenologicalFrameworkController.DeletePhenomenologicalFramework)

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

	// Language Optimization API (CRUD)
	r.POST("/api/language_optimization", LanguageOptimizationController.AddLanguageOptimization)
	r.GET("/api/language_optimization", LanguageOptimizationController.ListLanguageOptimizations)
	r.GET("/api/language_optimization/:id", LanguageOptimizationController.GetLanguageOptimization)
	r.PUT("/api/language_optimization/:id", LanguageOptimizationController.EditLanguageOptimization)
	r.DELETE("/api/language_optimization/:id", LanguageOptimizationController.DeleteLanguageOptimization)

	// Teaching Free Control API (CRUD)
	r.POST("/api/teaching_free_control", TeachingFreeControlController.AddTeachingFreeControl)
	r.GET("/api/teaching_free_control", TeachingFreeControlController.ListTeachingFreeControls)
	r.GET("/api/teaching_free_control/:id", TeachingFreeControlController.GetTeachingFreeControl)
	r.PUT("/api/teaching_free_control/:id", TeachingFreeControlController.EditTeachingFreeControl)
	r.DELETE("/api/teaching_free_control/:id", TeachingFreeControlController.DeleteTeachingFreeControl)

	// 404ハンドラー（一時的にコメントアウト）
	// r.NoRoute(middleware.NotFoundMiddleware())
	
	return r
}
