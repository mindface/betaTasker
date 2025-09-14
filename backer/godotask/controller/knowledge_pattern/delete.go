package knowledge_pattern

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteKnowledgePattern: DELETE /api/knowledge_pattern/:id
func (ctl *KnowledgePatternController) DeleteKnowledgePattern(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteKnowledgePattern(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete knowledge pattern"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Knowledge pattern deleted"})
}
