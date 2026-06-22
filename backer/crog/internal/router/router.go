package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
  r := gin.Default()

  r.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok!", "service": "crog"})
  })

  api := r.Group("/api/v1")
  {
     _ = api // ここにハンドラーを追加していく
  }

  return r
}
