package server

import (
	"github.com/godotask/controller/book"
	"github.com/godotask/controller/memory"
	"github.com/godotask/controller/top"
	"github.com/godotask/controller/user"
	"github.com/godotask/controller/task"
	"github.com/godotask/controller/assessment"
	"github.com/godotask/controller/heuristics"
	"github.com/godotask/controller/multimodal"
	// "github.com/godotask/controller/quantification_label"
	"github.com/godotask/controller/state_evaluation"
	"github.com/godotask/controller/tool_matching"
	"github.com/godotask/controller/process_monitoring"
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

	// Quantification Label and Multimodal Controllers
	quantificationLabelRepo := repository.NewQuantificationLabelRepository(model.DB)
	// Initialize database schema for quantification labels
	quantificationLabelRepo.InitializeDatabase()
	
	quantificationLabelService := service.NewQuantificationLabelService(model.DB)
	// quantificationLabelController := &quantification_label.QuantificationLabelController{
	// 	Service: quantificationLabelService,
	// }
	
	multimodalService := service.NewMultimodalService(model.DB)
	multimodalController := &multimodal.MultimodalController{
		Service:      multimodalService,
		LabelService: quantificationLabelService,
	}

	// State Evaluation, Tool Matching, and Process Monitoring Services
	stateEvaluationService := service.NewStateEvaluationService(model.DB)
	stateEvaluationController := &state_evaluation.StateEvaluationController{
		Service: stateEvaluationService,
	}
	
	toolMatchingService := service.NewToolMatchingService(model.DB)
	toolMatchingController := &tool_matching.ToolMatchingController{
		Service: toolMatchingService,
	}
	
	processMonitoringService := service.NewProcessMonitoringService(model.DB)
	processMonitoringController := &process_monitoring.ProcessMonitoringController{
		Service: processMonitoringService,
	}

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

	// Quantification Label API
	// r.GET("/api/labels", quantificationLabelController.GetLabels)
	// r.POST("/api/labels", quantificationLabelController.CreateLabel)
	// r.PUT("/api/labels/:id", quantificationLabelController.UpdateLabel)
	// r.DELETE("/api/labels/:id", quantificationLabelController.DeleteLabel)
	// r.GET("/api/labels/search", quantificationLabelController.SearchLabels)
	// r.POST("/api/labels/:id/verify", quantificationLabelController.VerifyLabel)
	// r.GET("/api/labels/statistics", quantificationLabelController.GetStatistics)
	// r.POST("/api/labels/suggest", quantificationLabelController.SuggestQuantification)
	// r.GET("/api/labels/:id/history", quantificationLabelController.GetLabelHistory)
	// r.GET("/api/labels/similar", quantificationLabelController.FindSimilar)
	// r.POST("/api/labels/bulk", quantificationLabelController.BulkOperation)
	// r.GET("/api/labels/export", quantificationLabelController.ExportLabels)
	// r.GET("/api/labels/user-stats", quantificationLabelController.GetUserStats)

	// Multimodal API
	r.POST("/api/multimodal/process", multimodalController.ProcessMultimodal)
	r.POST("/api/multimodal/upload", multimodalController.UploadImage)
	r.POST("/api/multimodal/calibrate", multimodalController.CalibrateUser)
	r.GET("/api/multimodal/calibration/:user_id", multimodalController.GetCalibration)
	r.POST("/api/multimodal/verify", multimodalController.VerifyQuantification)
	r.GET("/api/multimodal/export", multimodalController.ExportData)
	r.POST("/api/multimodal/compare", multimodalController.CompareImages)
	r.POST("/api/multimodal/feedback", multimodalController.SubmitFeedback)

	// State Evaluation API
	r.POST("/api/state-evaluations", stateEvaluationController.CreateEvaluation)
	r.GET("/api/state-evaluations/:id", stateEvaluationController.GetEvaluation)
	r.GET("/api/state-evaluations/user/:user_id", stateEvaluationController.GetEvaluationHistory)
	r.PUT("/api/state-evaluations/:id/results", stateEvaluationController.UpdateEvaluationResults)
	r.GET("/api/state-evaluations/user/:user_id/progression", stateEvaluationController.GetLevelProgression)
	r.GET("/api/state-evaluations/stats", stateEvaluationController.GetEvaluationStats)

	// Tool Matching API
	r.POST("/api/tool-matching", toolMatchingController.FindOptimalTools)
	r.GET("/api/tool-matching/state-evaluation/:state_evaluation_id", toolMatchingController.GetMatchingHistory)
	r.POST("/api/tool-matching/recommendations", toolMatchingController.GetRecommendations)
	r.GET("/api/tool-matching/available-tools", toolMatchingController.GetAvailableTools)
	r.POST("/api/tool-matching/compare", toolMatchingController.CompareTools)
	r.GET("/api/tool-matching/performance-prediction", toolMatchingController.PredictPerformance)

	// Process Monitoring API
	r.POST("/api/process-monitoring/start", processMonitoringController.StartMonitoring)
	r.POST("/api/process-monitoring/:id/stop", processMonitoringController.StopMonitoring)
	r.GET("/api/process-monitoring/:id/data", processMonitoringController.GetMonitoringData)
	r.GET("/api/process-monitoring/state-evaluation/:state_evaluation_id", processMonitoringController.GetMonitoringHistory)
	r.GET("/api/process-monitoring/active", processMonitoringController.GetActiveMonitors)
	r.GET("/api/process-monitoring/:id/status", processMonitoringController.GetMonitorStatus)
	r.GET("/api/process-monitoring/:id/ws", processMonitoringController.HandleWebSocket) // WebSocket endpoint
	r.GET("/api/process-monitoring/:id/summary", processMonitoringController.GetMonitoringSummary)
	r.POST("/api/process-monitoring/:id/alert-thresholds", processMonitoringController.UpdateAlertThresholds)

	// 404ハンドラー（一時的にコメントアウト）
	// r.NoRoute(middleware.NotFoundMiddleware())
	
	return r
}
