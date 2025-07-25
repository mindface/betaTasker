package memory

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// ListMemories: GET /api/memory
func (ctl *MemoryController) ListMemories(c *gin.Context) {
	memories, err := ctl.Service.ListMemories()
	fmt.Printf("ListMemories called, found %d memories\n", len(memories))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list memories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"memories": memories})
}
