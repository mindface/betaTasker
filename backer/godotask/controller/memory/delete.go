package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/service"
)

type MemoryController struct {
	Service *service.MemoryService
}

// DeleteMemory: DELETE /api/memory/:id
func (ctl *MemoryController) DeleteMemory(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteMemory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete memory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory deleted"})
}
