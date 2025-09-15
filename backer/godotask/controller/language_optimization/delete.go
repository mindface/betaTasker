package language_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteLanguageOptimization: DELETE /api/language_optimization/:id
func (ctl *LanguageOptimizationController) DeleteLanguageOptimization(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteLanguageOptimization(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete language optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Language optimization deleted"})
}
