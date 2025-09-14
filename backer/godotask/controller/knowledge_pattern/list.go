package knowledge_pattern

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListKnowledgePatterns: GET /api/knowledge_patterns
func (ctl *KnowledgePatternController) ListKnowledgePatterns(c *gin.Context) {
	knowledgePatterns, err := ctl.Service.ListKnowledgePatterns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list KnowledgePatterns"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"KnowledgePatterns": knowledgePatterns})
}