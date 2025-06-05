package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetMemory: GET /api/memory/:id
func (ctl *MemoryController) GetMemory(c *gin.Context) {
	id := c.Param("id")
	memory, err := ctl.Service.GetMemoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Memory not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"memory": memory})
}
