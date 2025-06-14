package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
	"fmt"
)

// AddMemory: POST /api/memory
func (ctl *MemoryController) AddMemory(c *gin.Context) {
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		fmt.Printf("AddMemory BindJSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateMemory(&memory); err != nil {
		fmt.Printf("AddMemory CreateMemory error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add memory"})
		return
	}
	fmt.Printf("AddMemory success: %+v\n", memory)
	c.JSON(http.StatusOK, gin.H{"message": "Memory added", "memory": memory})
}
