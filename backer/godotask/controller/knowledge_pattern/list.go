package knowledge_pattern

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// ListKnowledgePatterns: GET /api/knowledge_patterns
func (ctl *KnowledgePatternController) ListKnowledgePatterns(c *gin.Context) {
	knowledgePatterns, err := ctl.Service.ListKnowledgePatterns()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list knowledge patterns",
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
		"message": "Knowledge patterns retrieved",
		"KnowledgePatterns": knowledgePatterns,
	})
}