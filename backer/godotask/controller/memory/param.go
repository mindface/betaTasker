package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// GET /api/memory/context/:code
func (ctl *MemoryController) GetMemoryContextByCode(c *gin.Context) {
	code := c.Param("code") // 例: MA-Q-02
	var contexts []model.MemoryContext
	err := ctl.Service.FindMemoryContextsByCode(code, &contexts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get memory contexts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"contexts": contexts})
}
