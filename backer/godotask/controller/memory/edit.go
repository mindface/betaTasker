package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
	"github.com/godotask/service"
)

type MemoryController struct {
	Service *service.MemoryService
}

// EditMemory: PUT /api/memory/:id
func (ctl *MemoryController) EditMemory(c *gin.Context) {
	id := c.Param("id")
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateMemory(id, &memory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit memory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory edited", "memory": memory})
}
