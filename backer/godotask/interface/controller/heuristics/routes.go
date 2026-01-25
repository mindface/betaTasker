package heuristics

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes ヒューリスティクスルートの登録
func RegisterRoutes(router *gin.RouterGroup) {
	// MLパイプライン初期化
	pipeline := NewMLPipeline()
	
	// ヒューリスティクスグループ
	heuristics := router.Group("/heuristics")
	{
		// モデル同期
		heuristics.POST("/sync", pipeline.SyncHandler)
		
		// アクション記録
		heuristics.POST("/record", pipeline.RecordHandler)
		
		// モデル取得
		heuristics.GET("/model", pipeline.ModelHandler)
		
		// メトリクス取得
		heuristics.GET("/metrics", pipeline.MetricsHandler)
		
		// 予測
		heuristics.POST("/predict", pipeline.PredictHandler)
	}
}