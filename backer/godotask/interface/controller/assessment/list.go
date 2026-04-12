package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext"
	"github.com/godotask/interface/tools"
	"github.com/godotask/errors"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"

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
    pager := tools.ParsePagerQuery(c)
    filter := dtoquery.QueryFilter{
      UserID:  &pager.UserID,
      TaskID:  pager.TaskID,
      Include: helperquery.ParseIncludeParam(c.Query("include")),
    }
    // Service 側で total も返す想定
    assessments, total, err := ctl.Service.ListAssessmentsPager(*filter.UserID, pager)
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
      totalPages = int((total + int64(pager.Limit) - 1) / int64(pager.Limit))
    }

    c.JSON(http.StatusOK, gin.H{
      "success":     true,
      "message":     "Assessments retrieved",
      "assessments": assessments,
      "meta": gin.H{
        "total":       total,
        "total_pages": totalPages,
        "page":        pager.Page,
        "limit":       pager.Limit,
      },
    })
}

// ListAssessmentsForTaskUserPager: GET /api/assessment/task-user/pager
func (ctl *AssessmentController) ListAssessmentsForTaskUserPager(c *gin.Context) {
  // クエリパラメータ
  pager := tools.ParsePagerQuery(c)
  filter := dtoquery.QueryFilter{
    UserID: nil,
    TaskID:  pager.TaskID,
    Include: helperquery.ParseIncludeParam(c.Query("include")),
  }

  // Service 側で total も返す想定
  assessments, total, err := ctl.Service.ListAssessmentsForTaskUserPager(
    filter,pager,
  )

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
    "meta":    tools.BuildPageMeta(total, pager.Page, pager.Limit),
  })
}
