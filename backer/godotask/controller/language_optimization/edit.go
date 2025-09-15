package language_optimization

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditLanguageOptimization: PUT /api/language_optimization/:id
func (ctl *LanguageOptimizationController) EditLanguageOptimization(c *gin.Context) {
	id := c.Param("id")
	var languageOptimization model.LanguageOptimization
	if err := c.ShouldBindJSON(&languageOptimization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateLanguageOptimization(id, &languageOptimization); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit language optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Language optimization edited", "language_optimization": languageOptimization})
}
