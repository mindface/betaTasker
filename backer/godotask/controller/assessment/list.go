package assessment

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

type TaskUserRequest struct {
  TaskID int `json:"task_id"`
  UserID int `json:"user_id"`
}

// ListAssessments: GET /api/assessment
func (ctl *AssessmentController) ListAssessments(c *gin.Context) {
	assessments, err := ctl.Service.ListAssessments()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list assessments",
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
		"message": "Assessments retrieved",
		"assessments": assessments,
	})
}

func (ctl *AssessmentController) ListAssessmentsForTaskUser(c *gin.Context) {
	var req TaskUserRequest
	// JSON Body をバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Invalid request body",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	fmt.Printf("ListAssessmentsForTaskUser called with user_id: %d, task_id: %d\n", req.UserID, req.TaskID)

	assessments, err := ctl.Service.ListAssessmentsForTaskUser(req.UserID, req.TaskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assessments not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"assessments": assessments})
}
