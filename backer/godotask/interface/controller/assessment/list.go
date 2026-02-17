package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext"
	"github.com/godotask/interface/tools"
	"github.com/godotask/errors"
	"strconv"

	// "github.com/rs/zerolog/log"
)

type TaskUserRequest struct {
  TaskID int `json:"task_id"`
  UserID int `json:"user_id"`
}

// ListAssessments: GET /api/assessment
func (ctl *AssessmentController) ListAssessments(c *gin.Context) {

	userID, _ := authcontext.UserID(c)
	assessments, total, err := ctl.Service.ListAssessments(userID)
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
    "total": total,
	})
}

// ListAssessmentsPager: GET /api/assessment/pager
func (ctl *AssessmentController) ListAssessmentsPager(c *gin.Context) {
    // クエリパラメータ
    page := 1
    limit := 20
    const maxPerPage = 100
	  userID, _ := authcontext.UserID(c)

    if p := c.Query("page"); p != "" {
      if v, err := strconv.Atoi(p); err == nil && v > 0 {
        page = v
      }
    }
    if pp := c.Query("limit"); pp != "" {
      if v, err := strconv.Atoi(pp); err == nil && v > 0 {
        limit = v
      }
    }
    if limit > maxPerPage {
      limit = maxPerPage
    }

    offset := (page - 1) * limit

    // Service 側で total も返す想定
    assessments, total, err := ctl.Service.ListAssessmentPager(userID, limit, offset)
    if err != nil {
      appErr := errors.NewAppError(
        errors.SYS_INTERNAL_ERROR,
        errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
        err.Error(),
      )
      c.JSON(appErr.HTTPStatus, gin.H{
        "code":    appErr.Code,
        "message": appErr.Message,
        "detail":  appErr.Detail,
      })
      return
    }

    totalPages := 0
    if total > 0 {
      totalPages = int((total + int64(limit) - 1) / int64(limit))
    }

    c.JSON(http.StatusOK, gin.H{
      "success":     true,
      "message":     "Assessments retrieved",
      "assessments": assessments,
      "meta": gin.H{
        "total":       total,
        "total_pages": totalPages,
        "page":        page,
        "limit":       limit,
      },
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

	assessments, err := ctl.Service.ListAssessmentsForTaskUser(req.UserID, req.TaskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assessments not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"assessments": assessments})
}

// ListAssessmentsForTaskUserPager: GET /api/assessment/task-user/pager
func (ctl *AssessmentController) ListAssessmentsForTaskUserPager(c *gin.Context) {
  // クエリパラメータ
  query := tools.ParsePagerQuery(c)

  // Service 側で total も返す想定
  assessments, total, err := ctl.Service.ListAssessmentsForTaskUserPager(query.UserID, query.TaskID, query.Offset, query.Limit)
  if err != nil {
    appErr := errors.NewAppError(
      errors.SYS_INTERNAL_ERROR,
      errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
      err.Error(),
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
    "message": "assessments retrieved",
    "assessments": assessments,
    "meta":    tools.BuildPageMeta(total, query.Page, query.Limit),
  })
}
