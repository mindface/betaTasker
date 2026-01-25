package knowledge_pattern

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// GetKnowledgePattern: GET /api/knowledge_pattern/:id
func (ctl *KnowledgePatternController) GetKnowledgePattern(c *gin.Context) {
	id := c.Param("id")
	knowledgePattern, err := ctl.Service.GetKnowledgePatternByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Knowledge pattern not found",
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
		"message": "Knowledge pattern retrieved",
		"knowledge_pattern": knowledgePattern,
	})
}
