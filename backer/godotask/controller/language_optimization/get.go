package language_optimization

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetLanguageOptimization: GET /api/language_optimization/:id
func (ctl *LanguageOptimizationController) GetLanguageOptimization(c *gin.Context) {
	id := c.Param("id")
	languageOptimization, err := ctl.Service.GetLanguageOptimizationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language optimization not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"language_optimization": languageOptimization})
}
