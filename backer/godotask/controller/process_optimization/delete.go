package process_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteProcessOptimization: DELETE /api/process_optimization/:id
func (ctl *ProcessOptimizationController) DeleteProcessOptimization(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteProcessOptimization(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete process optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Process optimization deleted"})
}
