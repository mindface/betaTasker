package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetTask: GET /api/task/:id
func (ctl *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := ctl.Service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}
