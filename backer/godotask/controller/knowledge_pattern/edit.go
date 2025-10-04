package knowledge_pattern

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// EditKnowledgePattern: PUT /api/knowledge_pattern/:id
func (ctl *KnowledgePatternController) EditKnowledgePattern(c *gin.Context) {
	id := c.Param("id")
	var knowledgePattern model.KnowledgePattern
	if err := c.ShouldBindJSON(&knowledgePattern); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&knowledgePattern); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	if err := ctl.Service.UpdateKnowledgePattern(id, &knowledgePattern); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to edit process optimization",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Knowledge pattern edited",
		"knowledge_pattern": knowledgePattern,
	})
}
