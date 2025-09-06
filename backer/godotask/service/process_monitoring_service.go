package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProcessMonitoringService struct {
	db          *gorm.DB
	connections map[string]*websocket.Conn
	monitors    map[string]*ProcessMonitor
	mutex       sync.RWMutex
	upgrader    websocket.Upgrader
}

type ProcessMonitor struct {
	ID                string
	StateEvaluationID string
	ProcessType       string
	Status            string
	StartTime         time.Time
	LastUpdate        time.Time
	Metrics           map[string]interface{}
	Anomalies         []Anomaly
	Thresholds        map[string]float64
	TargetPerformance map[string]interface{}
	RealTimeData      chan MonitoringData
	StopChannel       chan bool
}

type Anomaly struct {
	Type        string    `json:"type"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	Threshold   float64   `json:"threshold"`
	Timestamp   time.Time `json:"timestamp"`
}

type MonitoringData struct {
	Timestamp   time.Time              `json:"timestamp"`
	ProcessID   string                 `json:"process_id"`
	Metrics     map[string]interface{} `json:"metrics"`
	Status      string                 `json:"status"`
	Anomalies   []Anomaly              `json:"anomalies"`
	Performance map[string]interface{} `json:"performance"`
}

func NewProcessMonitoringService(db *gorm.DB) *ProcessMonitoringService {
	return &ProcessMonitoringService{
		db:          db,
		connections: make(map[string]*websocket.Conn),
		monitors:    make(map[string]*ProcessMonitor),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *gin.Request) bool {
				return true // Allow connections from any origin in development
			},
		},
	}
}

func (s *ProcessMonitoringService) StartMonitoring(req *model.ProcessMonitoringRequest) (*model.ProcessMonitoring, error) {
	// Create process monitoring record
	monitoring := &model.ProcessMonitoring{
		ID:                uuid.New().String(),
		StateEvaluationID: req.StateEvaluationID,
		ProcessType:       req.ProcessType,
		Status:            "running",
		StartTime:         time.Now(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// Set initial monitoring data
	if req.InitialData != nil {
		initialDataJSON, err := json.Marshal(req.InitialData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal initial data: %v", err)
		}
		monitoring.MonitoringData = datatypes.JSON(initialDataJSON)
	}

	// Get state evaluation for context
	var stateEval model.StateEvaluation
	if err := s.db.Where("id = ?", req.StateEvaluationID).First(&stateEval).Error; err != nil {
		return nil, fmt.Errorf("state evaluation not found: %v", err)
	}

	// Get tool matching result for performance expectations
	var toolResult model.ToolMatchingResult
	if err := s.db.Where("state_evaluation_id = ?", req.StateEvaluationID).First(&toolResult).Error; err != nil {
		log.Printf("No tool matching result found for state evaluation %s", req.StateEvaluationID)
	}

	// Create process monitor
	monitor := &ProcessMonitor{
		ID:                monitoring.ID,
		StateEvaluationID: req.StateEvaluationID,
		ProcessType:       req.ProcessType,
		Status:            "running",
		StartTime:         time.Now(),
		LastUpdate:        time.Now(),
		Metrics:           make(map[string]interface{}),
		Anomalies:         []Anomaly{},
		RealTimeData:      make(chan MonitoringData, 100),
		StopChannel:       make(chan bool, 1),
	}

	// Set thresholds based on process type and state evaluation
	monitor.Thresholds = s.generateThresholds(req.ProcessType, &stateEval)

	// Set target performance from tool matching result
	if toolResult.ID != "" {
		var expectedPerformance map[string]interface{}
		if err := json.Unmarshal(toolResult.ExpectedPerformance, &expectedPerformance); err == nil {
			monitor.TargetPerformance = expectedPerformance
		}
	}

	// Store monitor
	s.mutex.Lock()
	s.monitors[monitoring.ID] = monitor
	s.mutex.Unlock()

	// Start monitoring goroutine
	go s.runMonitoring(monitor)

	// Save to database
	if err := s.db.Create(monitoring).Error; err != nil {
		return nil, fmt.Errorf("failed to save monitoring record: %v", err)
	}

	return monitoring, nil
}

func (s *ProcessMonitoringService) generateThresholds(processType string, stateEval *model.StateEvaluation) map[string]float64 {
	baseThresholds := map[string]map[string]float64{
		"robot_assembly": {
			"force_variation":     10.0, // N
			"position_error":      0.05, // mm
			"cycle_time_increase": 1.2,  // 20% increase threshold
			"success_rate_drop":   0.05, // 5% drop threshold
		},
		"robot_welding": {
			"current_variation":   20.0, // A
			"voltage_variation":   5.0,  // V
			"speed_variation":     2.0,  // mm/s
			"quality_drop":        0.1,  // 10% quality drop
		},
		"machining": {
			"vibration_increase":  2.0,  // 2x increase
			"temperature_rise":    20.0, // °C
			"surface_roughness":   0.5,  // μm increase
			"tool_wear_rate":      0.1,  // mm/hour
		},
		"robot_vision": {
			"recognition_drop":    0.05, // 5% drop
			"processing_time":     50.0, // ms increase
			"false_positive_rate": 0.02, // 2% increase
			"lighting_variation":  20.0, // % variation
		},
		"robot_motion": {
			"trajectory_error":    1.0,  // mm
			"velocity_variation":  100.0, // mm/s
			"acceleration_spike":  500.0, // mm/s²
			"jerk_limit":          1000.0, // mm/s³
		},
	}

	thresholds := baseThresholds[processType]
	if thresholds == nil {
		// Default thresholds
		thresholds = map[string]float64{
			"performance_drop":    0.1,  // 10% performance drop
			"error_rate_increase": 0.05, // 5% error rate increase
			"response_time":       100.0, // ms
			"resource_usage":      80.0, // % usage
		}
	}

	// Adjust thresholds based on user level
	levelMultiplier := []float64{1.5, 1.3, 1.0, 0.8, 0.6}[stateEval.Level-1]
	
	adjustedThresholds := make(map[string]float64)
	for key, value := range thresholds {
		adjustedThresholds[key] = value * levelMultiplier
	}

	return adjustedThresholds
}

func (s *ProcessMonitoringService) runMonitoring(monitor *ProcessMonitor) {
	ticker := time.NewTicker(1 * time.Second) // Update every second
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate simulated monitoring data
			data := s.generateMonitoringData(monitor)
			
			// Check for anomalies
			anomalies := s.detectAnomalies(monitor, data.Metrics)
			data.Anomalies = anomalies
			
			// Update monitor
			monitor.LastUpdate = time.Now()
			monitor.Metrics = data.Metrics
			monitor.Anomalies = append(monitor.Anomalies, anomalies...)
			
			// Limit anomaly history
			if len(monitor.Anomalies) > 100 {
				monitor.Anomalies = monitor.Anomalies[len(monitor.Anomalies)-100:]
			}
			
			// Send to WebSocket clients
			select {
			case monitor.RealTimeData <- data:
			default:
				// Channel full, drop oldest data
			}
			
			// Broadcast to WebSocket connections
			s.broadcastToConnections(monitor.ID, data)
			
			// Update database periodically (every 10 seconds)
			if time.Since(monitor.LastUpdate).Seconds() >= 10 {
				s.updateMonitoringRecord(monitor)
			}

		case <-monitor.StopChannel:
			// Stop monitoring
			monitor.Status = "stopped"
			s.updateMonitoringRecord(monitor)
			return
		}
	}
}

func (s *ProcessMonitoringService) generateMonitoringData(monitor *ProcessMonitor) MonitoringData {
	data := MonitoringData{
		Timestamp: time.Now(),
		ProcessID: monitor.ID,
		Status:    monitor.Status,
	}

	// Generate metrics based on process type
	switch monitor.ProcessType {
	case "robot_assembly":
		data.Metrics = s.generateAssemblyMetrics(monitor)
	case "robot_welding":
		data.Metrics = s.generateWeldingMetrics(monitor)
	case "machining":
		data.Metrics = s.generateMachiningMetrics(monitor)
	case "robot_vision":
		data.Metrics = s.generateVisionMetrics(monitor)
	case "robot_motion":
		data.Metrics = s.generateMotionMetrics(monitor)
	default:
		data.Metrics = s.generateGenericMetrics(monitor)
	}

	// Calculate performance metrics
	data.Performance = s.calculatePerformanceMetrics(monitor, data.Metrics)

	return data
}

func (s *ProcessMonitoringService) generateAssemblyMetrics(monitor *ProcessMonitor) map[string]interface{} {
	// Simulate assembly process metrics
	baseTime := time.Since(monitor.StartTime).Seconds()
	
	return map[string]interface{}{
		"force_x":           5.0 + 3.0*math.Sin(baseTime/10) + rand.Float64()*2-1,
		"force_y":           3.0 + 2.0*math.Cos(baseTime/8) + rand.Float64()*1.5-0.75,
		"force_z":           15.0 + 5.0*math.Sin(baseTime/15) + rand.Float64()*3-1.5,
		"position_error":    0.01 + rand.Float64()*0.02,
		"cycle_time":        25.0 + rand.Float64()*5-2.5,
		"success_rate":      0.95 + rand.Float64()*0.05,
		"temperature":       35.0 + rand.Float64()*10,
		"parts_assembled":   int(baseTime / 30), // One part every 30 seconds
	}
}

func (s *ProcessMonitoringService) generateWeldingMetrics(monitor *ProcessMonitor) map[string]interface{} {
	baseTime := time.Since(monitor.StartTime).Seconds()
	
	return map[string]interface{}{
		"current":           200.0 + 50.0*math.Sin(baseTime/5) + rand.Float64()*20-10,
		"voltage":           25.0 + 3.0*math.Cos(baseTime/7) + rand.Float64()*2-1,
		"speed":             10.0 + rand.Float64()*2-1,
		"penetration":       3.5 + rand.Float64()*0.5-0.25,
		"quality_score":     0.92 + rand.Float64()*0.08,
		"arc_stability":     0.88 + rand.Float64()*0.1,
		"spatter_count":     rand.Intn(5),
		"weld_length":       baseTime * 10, // mm welded
	}
}

func (s *ProcessMonitoringService) generateMachiningMetrics(monitor *ProcessMonitor) map[string]interface{} {
	baseTime := time.Since(monitor.StartTime).Seconds()
	
	return map[string]interface{}{
		"spindle_speed":     2000.0 + 500.0*math.Sin(baseTime/20) + rand.Float64()*100-50,
		"feed_rate":         150.0 + rand.Float64()*20-10,
		"cutting_force":     80.0 + 20.0*math.Sin(baseTime/8) + rand.Float64()*15-7.5,
		"vibration":         0.5 + 0.2*math.Sin(baseTime/12) + rand.Float64()*0.1-0.05,
		"temperature":       45.0 + rand.Float64()*15,
		"surface_roughness": 1.2 + rand.Float64()*0.4-0.2,
		"tool_wear":         baseTime * 0.001, // mm per second
		"chips_evacuated":   rand.Float64() > 0.1, // 90% success rate
	}
}

func (s *ProcessMonitoringService) generateVisionMetrics(monitor *ProcessMonitor) map[string]interface{} {
	baseTime := time.Since(monitor.StartTime).Seconds()
	
	return map[string]interface{}{
		"recognition_accuracy": 0.96 + rand.Float64()*0.04-0.02,
		"processing_time":      25.0 + rand.Float64()*10-5,
		"objects_detected":     rand.Intn(5) + 1,
		"false_positives":      rand.Intn(2),
		"lighting_level":       80.0 + 20.0*math.Sin(baseTime/30) + rand.Float64()*10-5,
		"image_quality":        0.85 + rand.Float64()*0.15,
		"fps":                  30.0 + rand.Float64()*5-2.5,
		"memory_usage":         65.0 + rand.Float64()*20,
	}
}

func (s *ProcessMonitoringService) generateMotionMetrics(monitor *ProcessMonitor) map[string]interface{} {
	baseTime := time.Since(monitor.StartTime).Seconds()
	
	return map[string]interface{}{
		"position_x":        100.0 + 50.0*math.Sin(baseTime/10),
		"position_y":        200.0 + 30.0*math.Cos(baseTime/8),
		"position_z":        150.0 + 20.0*math.Sin(baseTime/12),
		"velocity":          500.0 + 100.0*math.Sin(baseTime/5) + rand.Float64()*50-25,
		"acceleration":      100.0 + rand.Float64()*50-25,
		"jerk":              50.0 + rand.Float64()*30-15,
		"path_error":        0.05 + rand.Float64()*0.03,
		"joint_angles":      []float64{0.5, 1.2, -0.8, 2.1, 0.3, -1.5},
	}
}

func (s *ProcessMonitoringService) generateGenericMetrics(monitor *ProcessMonitor) map[string]interface{} {
	return map[string]interface{}{
		"performance_score": 0.85 + rand.Float64()*0.15,
		"error_rate":        0.02 + rand.Float64()*0.03,
		"response_time":     50.0 + rand.Float64()*30,
		"resource_usage":    60.0 + rand.Float64()*25,
		"throughput":        100.0 + rand.Float64()*20-10,
		"availability":      0.98 + rand.Float64()*0.02,
	}
}

func (s *ProcessMonitoringService) calculatePerformanceMetrics(monitor *ProcessMonitor, metrics map[string]interface{}) map[string]interface{} {
	performance := make(map[string]interface{})

	// Calculate overall efficiency
	efficiency := s.calculateEfficiency(monitor, metrics)
	performance["efficiency"] = efficiency

	// Calculate quality score
	quality := s.calculateQuality(monitor, metrics)
	performance["quality"] = quality

	// Calculate stability score
	stability := s.calculateStability(monitor, metrics)
	performance["stability"] = stability

	// Overall performance score
	overall := (efficiency + quality + stability) / 3.0
	performance["overall"] = overall

	// Time-based metrics
	runtime := time.Since(monitor.StartTime)
	performance["runtime_seconds"] = runtime.Seconds()
	performance["uptime_percentage"] = s.calculateUptime(monitor)

	return performance
}

func (s *ProcessMonitoringService) calculateEfficiency(monitor *ProcessMonitor, metrics map[string]interface{}) float64 {
	switch monitor.ProcessType {
	case "robot_assembly":
		if cycleTime, ok := metrics["cycle_time"].(float64); ok {
			targetCycleTime := 25.0 // seconds
			return math.Max(0.0, math.Min(1.0, targetCycleTime/cycleTime))
		}
	case "robot_welding":
		if quality, ok := metrics["quality_score"].(float64); ok {
			return quality
		}
	case "machining":
		if surfaceRoughness, ok := metrics["surface_roughness"].(float64); ok {
			targetRoughness := 1.6
			return math.Max(0.0, math.Min(1.0, targetRoughness/surfaceRoughness))
		}
	}
	return 0.85 + rand.Float64()*0.1 // Default random efficiency
}

func (s *ProcessMonitoringService) calculateQuality(monitor *ProcessMonitor, metrics map[string]interface{}) float64 {
	switch monitor.ProcessType {
	case "robot_assembly":
		if successRate, ok := metrics["success_rate"].(float64); ok {
			return successRate
		}
	case "robot_vision":
		if accuracy, ok := metrics["recognition_accuracy"].(float64); ok {
			return accuracy
		}
	}
	return 0.90 + rand.Float64()*0.08 // Default random quality
}

func (s *ProcessMonitoringService) calculateStability(monitor *ProcessMonitor, metrics map[string]interface{}) float64 {
	// Calculate stability based on variation in metrics
	variationScore := 1.0
	
	// Simple stability calculation - in practice, this would analyze historical variance
	anomalyCount := len(monitor.Anomalies)
	if anomalyCount > 0 {
		recentAnomalies := 0
		cutoff := time.Now().Add(-5 * time.Minute)
		for _, anomaly := range monitor.Anomalies {
			if anomaly.Timestamp.After(cutoff) {
				recentAnomalies++
			}
		}
		variationScore = math.Max(0.0, 1.0-float64(recentAnomalies)*0.1)
	}
	
	return variationScore
}

func (s *ProcessMonitoringService) calculateUptime(monitor *ProcessMonitor) float64 {
	// Simplified uptime calculation - would be more complex in practice
	return 0.99 + rand.Float64()*0.01
}

func (s *ProcessMonitoringService) detectAnomalies(monitor *ProcessMonitor, metrics map[string]interface{}) []Anomaly {
	var anomalies []Anomaly

	for metricName, threshold := range monitor.Thresholds {
		if value, exists := metrics[metricName]; exists {
			if floatValue, ok := value.(float64); ok {
				if math.Abs(floatValue) > threshold {
					severity := "medium"
					if math.Abs(floatValue) > threshold*1.5 {
						severity = "high"
					}

					anomaly := Anomaly{
						Type:        metricName,
						Severity:    severity,
						Description: fmt.Sprintf("%s value %.2f exceeds threshold %.2f", metricName, floatValue, threshold),
						Value:       floatValue,
						Threshold:   threshold,
						Timestamp:   time.Now(),
					}
					anomalies = append(anomalies, anomaly)
				}
			}
		}
	}

	return anomalies
}

func (s *ProcessMonitoringService) broadcastToConnections(monitorID string, data MonitoringData) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	message, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling monitoring data: %v", err)
		return
	}

	for connID, conn := range s.connections {
		// Check if this connection is interested in this monitor
		if strings.Contains(connID, monitorID) {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Error sending WebSocket message: %v", err)
				// Remove broken connection
				delete(s.connections, connID)
				conn.Close()
			}
		}
	}
}

func (s *ProcessMonitoringService) updateMonitoringRecord(monitor *ProcessMonitor) {
	metricsJSON, err := json.Marshal(monitor.Metrics)
	if err != nil {
		log.Printf("Error marshaling metrics: %v", err)
		return
	}

	anomaliesJSON, err := json.Marshal(monitor.Anomalies)
	if err != nil {
		log.Printf("Error marshaling anomalies: %v", err)
		return
	}

	updates := map[string]interface{}{
		"monitoring_data": datatypes.JSON(metricsJSON),
		"anomalies":       datatypes.JSON(anomaliesJSON),
		"status":          monitor.Status,
		"updated_at":      time.Now(),
	}

	if monitor.Status == "stopped" {
		updates["end_time"] = time.Now()
	}

	if err := s.db.Model(&model.ProcessMonitoring{}).Where("id = ?", monitor.ID).Updates(updates).Error; err != nil {
		log.Printf("Error updating monitoring record: %v", err)
	}
}

func (s *ProcessMonitoringService) StopMonitoring(monitorID string) error {
	s.mutex.RLock()
	monitor, exists := s.monitors[monitorID]
	s.mutex.RUnlock()

	if !exists {
		return fmt.Errorf("monitor not found: %s", monitorID)
	}

	// Signal stop
	select {
	case monitor.StopChannel <- true:
	default:
		// Channel might be full or closed
	}

	// Remove from active monitors
	s.mutex.Lock()
	delete(s.monitors, monitorID)
	s.mutex.Unlock()

	return nil
}

func (s *ProcessMonitoringService) HandleWebSocket(c *gin.Context) {
	monitorID := c.Param("id")
	if monitorID == "" {
		c.JSON(400, gin.H{"error": "Monitor ID required"})
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	// Generate connection ID
	connID := fmt.Sprintf("%s_%s", monitorID, uuid.New().String())

	// Store connection
	s.mutex.Lock()
	s.connections[connID] = conn
	s.mutex.Unlock()

	// Clean up on disconnect
	defer func() {
		s.mutex.Lock()
		delete(s.connections, connID)
		s.mutex.Unlock()
	}()

	// Check if monitor exists
	s.mutex.RLock()
	monitor, exists := s.monitors[monitorID]
	s.mutex.RUnlock()

	if !exists {
		conn.WriteMessage(websocket.TextMessage, []byte(`{"error": "Monitor not found"}`))
		return
	}

	// Send initial data
	if len(monitor.Metrics) > 0 {
		initialData := MonitoringData{
			Timestamp:   time.Now(),
			ProcessID:   monitor.ID,
			Metrics:     monitor.Metrics,
			Status:      monitor.Status,
			Anomalies:   monitor.Anomalies,
			Performance: s.calculatePerformanceMetrics(monitor, monitor.Metrics),
		}

		if message, err := json.Marshal(initialData); err == nil {
			conn.WriteMessage(websocket.TextMessage, message)
		}
	}

	// Keep connection alive and handle client messages
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		// Echo client messages (can be extended for client commands)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}
}

func (s *ProcessMonitoringService) GetMonitoringData(monitorID string, limit int) ([]MonitoringData, error) {
	s.mutex.RLock()
	monitor, exists := s.monitors[monitorID]
	s.mutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("monitor not found: %s", monitorID)
	}

	var data []MonitoringData
	
	// Collect data from channel (non-blocking)
	collected := 0
	for {
		select {
		case d := <-monitor.RealTimeData:
			data = append(data, d)
			collected++
			if limit > 0 && collected >= limit {
				return data, nil
			}
		default:
			// No more data in channel
			return data, nil
		}
	}
}

func (s *ProcessMonitoringService) GetMonitoringHistory(stateEvaluationID string) ([]model.ProcessMonitoring, error) {
	var monitoringRecords []model.ProcessMonitoring

	if err := s.db.Where("state_evaluation_id = ?", stateEvaluationID).Order("created_at DESC").Find(&monitoringRecords).Error; err != nil {
		return nil, fmt.Errorf("failed to get monitoring history: %v", err)
	}

	return monitoringRecords, nil
}

func (s *ProcessMonitoringService) GetActiveMonitors() map[string]*ProcessMonitor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Return copy to avoid race conditions
	activeMonitors := make(map[string]*ProcessMonitor)
	for id, monitor := range s.monitors {
		activeMonitors[id] = monitor
	}

	return activeMonitors
}