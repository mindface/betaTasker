package assessment

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type TaskUserRequest struct {
  TaskID int `json:"task_id"`
  UserID int `json:"user_id"`
}

// ListAssessments: GET /api/assessment
func (ctl *AssessmentController) ListAssessments(c *gin.Context) {
	assessments, err := ctl.Service.ListAssessments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list assessments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"assessments": assessments})
}

func (ctl *AssessmentController) ListAssessmentsForTaskUser(c *gin.Context) {
	var req TaskUserRequest
	// JSON Body をバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
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
