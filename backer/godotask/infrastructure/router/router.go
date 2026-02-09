package router

import (
	"github.com/godotask/interface/http/controller"
	"github.com/godotask/interface/controller/book"
	"github.com/godotask/interface/controller/memory"
	"github.com/godotask/interface/controller/task"
	"github.com/godotask/interface/controller/assessment"
	"github.com/godotask/interface/controller/heuristics"
	"github.com/godotask/interface/controller/heuristics/analyze"
	"github.com/godotask/interface/controller/heuristics/insight"
	"github.com/godotask/interface/controller/process_optimization"
	"github.com/godotask/interface/controller/qualitative_label"
	"github.com/godotask/interface/controller/knowledge_pattern"
	"github.com/godotask/interface/controller/language_optimization"
	"github.com/godotask/interface/controller/teaching_free_control"
	"github.com/godotask/interface/controller/phenomenological_framework"
	"github.com/godotask/usecase/service"
	"github.com/godotask/infrastructure/db/repository"
	"github.com/godotask/infrastructure/db/model"
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

func setupRouter() *gin.Engine {
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
	public := r.Group("/api")
	{
		public.POST("/login", authController.Login)
		public.POST("/register", authController.Register)
		public.POST("/logout", authController.Logout)
	}

	// Book API (CRUD)
	r.GET("/book", bookController.ListBooks)
	r.POST("/file", book.HundleUplond)
	r.POST("/book", bookController.AddBook)
	r.DELETE("/deletebook/:id", bookController.DeleteBook)
	r.PUT("/updatebook/:id", bookController.EditBook)

	// ===== 認証必須のエンドポイント =====
	// AuthMiddleware をこのグループに一括適用
	protected := r.Group("/api")
	protected.Use(authMiddleware)
	{
		// Task API (CRUD)
		protected.POST("/task", taskController.AddTask)
		protected.GET("/task", taskController.ListTasks)
		protected.GET("/task/pager", taskController.ListTasksPager)
		protected.GET("/task/:id", taskController.GetTask)
		protected.PUT("/task/:id", taskController.EditTask)
		protected.DELETE("/task/:id", taskController.DeleteTask)

		// User profile
		protected.GET("/user/profile", controller.Profile)

		// Memory API (CRUD)
		protected.POST("/memory", memoryController.AddMemory)
		protected.GET("/memory", memoryController.ListMemories)
		protected.GET("/memory/pager", memoryController.ListMemoriesPager)
		protected.GET("/memory/:id", memoryController.GetMemory)
		protected.GET("/memory/context/:code", memoryController.GetMemoryContextByCode)
		protected.GET("/memory/aid/:code", memoryController.GetMemoryAidByCode)
		protected.PUT("/memory/:id", memoryController.EditMemory)
		protected.DELETE("/memory/:id", memoryController.DeleteMemory)

		// Assessment API (CRUD)
		protected.POST("/assessment", assessmentController.AddAssessment)
		protected.GET("/assessment", assessmentController.ListAssessments)
		protected.GET("/assessment/pager", assessmentController.ListAssessmentsPager)
		protected.POST("/assessmentsForTaskUser", assessmentController.ListAssessmentsForTaskUser)
		protected.GET("/assessmentsForTaskUser/pager", assessmentController.ListAssessmentsForTaskUserPager)
		protected.GET("/assessment/:id", assessmentController.GetAssessment)
		protected.PUT("/assessment/:id", assessmentController.EditAssessment)
		protected.DELETE("/assessment/:id", assessmentController.DeleteAssessment)	

		// Heuristics API (ML Pipeline & Analytics)
		protected.POST("/heuristics/analyze", AnalyzeController.AddAnalyzeData)
		protected.GET("/heuristics/analyze/:id", AnalyzeController.GetAnalyzeData)
		protected.PUT("/heuristics/analyze/:id", AnalyzeController.EditAnalyzeData)
		protected.DELETE("/heuristics/analyze/:id", AnalyzeController.DeleteAnalyzeData)

		// Heuristics API (ML Pipeline & Analytics)
		protected.POST("/heuristics/insight", heuristicsInsightController.AddInsightData)
		protected.GET("/heuristics/insight/:id", heuristicsInsightController.GetInsightData)
		protected.PUT("/heuristics/insight/:id", heuristicsInsightController.EditInsightsData)
		protected.DELETE("/heuristics/insight/:id", heuristicsInsightController.DeleteInsightData)

		protected.POST("/heuristics/track", heuristicsController.TrackBehavior)
		protected.GET("/heuristics/track/:user_id", heuristicsController.GetTrackingData)
		protected.GET("/heuristics/patterns", heuristicsController.DetectPatterns)
		protected.POST("/heuristics/patterns/train", heuristicsController.TrainModel)

		// phenomenological framework API (CRUD)
		protected.POST("/phenomenological_framework", phenomenologicalFrameworkController.AddPhenomenologicalFramework)
		protected.GET("/phenomenological_framework", phenomenologicalFrameworkController.ListPhenomenologicalFrameworks)
		protected.GET("/phenomenological_framework/:id", phenomenologicalFrameworkController.GetPhenomenologicalFramework)
		protected.PUT("/phenomenological_framework/:id", phenomenologicalFrameworkController.EditPhenomenologicalFramework)
		protected.DELETE("/phenomenological_framework/:id", phenomenologicalFrameworkController.DeletePhenomenologicalFramework)

		// Process Optimization API (CRUD)
		protected.POST("/process_optimization", processOptimizationController.AddProcessOptimization)
		protected.GET("/process_optimization", processOptimizationController.ListProcessOptimizations)
		protected.GET("/process_optimization/:id", processOptimizationController.GetProcessOptimization)
		protected.PUT("/process_optimization/:id", processOptimizationController.EditProcessOptimization)
		protected.DELETE("/process_optimization/:id", processOptimizationController.DeleteProcessOptimization)

		// Qualitative Label API (CRUD)
		protected.POST("/qualitative_label", qualitativeLabelController.AddQualitativeLabel)
		protected.GET("/qualitative_label", qualitativeLabelController.ListQualitativeLabels)
		protected.GET("/qualitative_label/:id", qualitativeLabelController.GetQualitativeLabel)
		protected.PUT("/qualitative_label/:id", qualitativeLabelController.EditQualitativeLabel)
		protected.DELETE("/qualitative_label/:id", qualitativeLabelController.DeleteQualitativeLabel)

		// Knowledge Pattern API (CRUD)
		protected.POST("/knowledge_pattern", knowledgePatternController.AddKnowledgePattern)
		protected.GET("/knowledge_pattern", knowledgePatternController.ListKnowledgePatterns)
		protected.GET("/knowledge_pattern/:id", knowledgePatternController.GetKnowledgePattern)
		protected.PUT("/knowledge_pattern/:id", knowledgePatternController.EditKnowledgePattern)
		protected.DELETE("/knowledge_pattern/:id", knowledgePatternController.DeleteKnowledgePattern)

		// Language Optimization API (CRUD)
		protected.POST("/language_optimization", LanguageOptimizationController.AddLanguageOptimization)
		protected.GET("/language_optimization", LanguageOptimizationController.ListLanguageOptimizations)
		protected.GET("/language_optimization/:id", LanguageOptimizationController.GetLanguageOptimization)
		protected.PUT("/language_optimization/:id", LanguageOptimizationController.EditLanguageOptimization)
		protected.DELETE("/language_optimization/:id", LanguageOptimizationController.DeleteLanguageOptimization)

		// Teaching Free Control API (CRUD)
		protected.POST("/teaching_free_control", TeachingFreeControlController.AddTeachingFreeControl)
		protected.GET("/teaching_free_control", TeachingFreeControlController.ListTeachingFreeControls)
		protected.GET("/teaching_free_control/:id", TeachingFreeControlController.GetTeachingFreeControl)
		protected.PUT("/teaching_free_control/:id", TeachingFreeControlController.EditTeachingFreeControl)
		protected.DELETE("/teaching_free_control/:id", TeachingFreeControlController.DeleteTeachingFreeControl)
	}

	// 404ハンドラー（一時的にコメントアウト）
	// r.NoRoute(middleware.NotFoundMiddleware())
	
	return r
}

func GetRouter() *gin.Engine {
	return router
}
