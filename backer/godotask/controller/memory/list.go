package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/service"
)

type MemoryController struct {
	Service *service.MemoryService
}

// ListMemories: GET /api/memory
func (ctl *MemoryController) ListMemories(c *gin.Context) {
	memories, err := ctl.Service.ListMemories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list memories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"memories": memories})
}
