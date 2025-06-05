package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

// AddTask: POST /api/task
func (ctl *TaskController) AddTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task added", "task": task})
}
