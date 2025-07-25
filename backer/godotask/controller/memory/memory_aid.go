package memory

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// GET /api/memory/aid/:code
func (ctl *MemoryController) GetMemoryAidByCode(c *gin.Context) {
	code := c.Param("code")
	fmt.Printf("GetMemoryAidByCode code: %s", code)
	var contexts []model.MemoryContext
	err := ctl.Service.FindMemoryAidsByCode(code, &contexts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get memory aids"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"contexts": contexts})
}
