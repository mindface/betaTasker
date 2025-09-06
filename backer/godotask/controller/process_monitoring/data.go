package process_monitoring

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMonitoringData: GET /api/process-monitoring/:id/data
func (ctrl *ProcessMonitoringController) GetMonitoringData(c *gin.Context) {
	monitorID := c.Param("id")
	if monitorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Monitor ID is required"})
		return
	}

	limitStr := c.Query("limit")
	limit := 100 // Default limit
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	data, err := ctrl.Service.GetMonitoringData(monitorID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
		"count":  len(data),
	})
}