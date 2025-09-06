package process_monitoring

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/service"
)

// GetMonitoringHistory: GET /api/process-monitoring/state-evaluation/:state_evaluation_id
func (ctrl *ProcessMonitoringController) GetMonitoringHistory(c *gin.Context) {
	stateEvaluationID := c.Param("state_evaluation_id")
	if stateEvaluationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State evaluation ID is required"})
		return
	}

	history, err := ctrl.Service.GetMonitoringHistory(stateEvaluationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   history,
		"count":  len(history),
	})
}

// GetActiveMonitors: GET /api/process-monitoring/active
func (ctrl *ProcessMonitoringController) GetActiveMonitors(c *gin.Context) {
	activeMonitors := ctrl.Service.GetActiveMonitors()

	// Convert to response format (without exposing internal channels)
	response := make(map[string]gin.H)
	for id, monitor := range activeMonitors {
		response[id] = gin.H{
			"id":                   monitor.ID,
			"state_evaluation_id":  monitor.StateEvaluationID,
			"process_type":         monitor.ProcessType,
			"status":               monitor.Status,
			"start_time":           monitor.StartTime,
			"last_update":          monitor.LastUpdate,
			"metrics":              monitor.Metrics,
			"anomaly_count":        len(monitor.Anomalies),
			"recent_anomalies":     ctrl.getRecentAnomalies(monitor.Anomalies, 5),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
		"count":  len(response),
	})
}

// Helper function to get recent anomalies
func (ctrl *ProcessMonitoringController) getRecentAnomalies(anomalies []service.Anomaly, limit int) []service.Anomaly {
	if len(anomalies) <= limit {
		return anomalies
	}
	return anomalies[len(anomalies)-limit:]
}