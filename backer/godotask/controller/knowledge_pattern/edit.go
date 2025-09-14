package knowledge_pattern

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditKnowledgePattern: PUT /api/knowledge_pattern/:id
func (ctl *KnowledgePatternController) EditKnowledgePattern(c *gin.Context) {
	id := c.Param("id")
	var KnowledgePattern model.KnowledgePattern
	if err := c.ShouldBindJSON(&KnowledgePattern); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateKnowledgePattern(id, &KnowledgePattern); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit process optimization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Knowledge pattern edited", "knowledge_pattern": KnowledgePattern})
}
