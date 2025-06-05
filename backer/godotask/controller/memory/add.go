package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// AddMemory: POST /api/memory
func (ctl *MemoryController) AddMemory(c *gin.Context) {
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateMemory(&memory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add memory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory added", "memory": memory})
}
