package language_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListLanguageOptimizations: GET /api/language_optimization
func (ctl *LanguageOptimizationController) ListLanguageOptimizations(c *gin.Context) {
	languageOptimizations, err := ctl.Service.ListLanguageOptimizations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list LanguageOptimizations"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"LanguageOptimizations": languageOptimizations})
}