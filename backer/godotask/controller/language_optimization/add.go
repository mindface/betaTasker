package language_optimization

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddLanguageOptimization: POST /api/language_optimization
func (ctl *LanguageOptimizationController) AddLanguageOptimization(c *gin.Context) {
	var languageOptimization model.LanguageOptimization
	if err := c.ShouldBindJSON(&languageOptimization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateLanguageOptimization(&languageOptimization); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add language optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Language optimization added", "language_optimization": languageOptimization})
}
