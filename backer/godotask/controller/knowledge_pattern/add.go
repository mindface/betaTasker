package knowledge_pattern

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// AddKnowledgePattern: POST /api/knowledge_pattern
func (ctl *KnowledgePatternController) AddKnowledgePattern(c *gin.Context) {
	var knowledgePattern model.KnowledgePattern
	if err := c.ShouldBindJSON(&knowledgePattern); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateKnowledgePattern(&knowledgePattern); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add knowledge pattern"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Knowledge pattern added", "knowledge_pattern": knowledgePattern})
}
