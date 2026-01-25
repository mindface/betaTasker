package knowledge_pattern

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteKnowledgePattern: DELETE /api/knowledge_pattern/:id
func (ctl *KnowledgePatternController) DeleteKnowledgePattern(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteKnowledgePattern(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete knowledge pattern",
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
		"message": "Knowledge pattern deleted",
		"knowledge_pattern_id": id,
	})
}
