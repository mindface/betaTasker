package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ListTasks: GET /api/task
func (ctl *TaskController) ListTasks(c *gin.Context) {

	tasks, err := ctl.Service.ListTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list tasks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
