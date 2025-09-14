package knowledge_pattern

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetKnowledgePattern: GET /api/knowledge_pattern/:id
func (ctl *KnowledgePatternController) GetKnowledgePattern(c *gin.Context) {
	id := c.Param("id")
	knowledgePattern, err := ctl.Service.GetKnowledgePatternByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Knowledge pattern not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"knowledge_pattern": knowledgePattern})
}
