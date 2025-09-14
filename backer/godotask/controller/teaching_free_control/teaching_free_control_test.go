package teaching_free_control

import (
  "bytes"
  "net/http"
  "net/http/httptest"
  "testing"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  "github.com/godotask/model"
  "github.com/godotask/service"
)

// モックリポジトリ
type MockTeachingFreeControlRepository struct{}

func (m *MockTeachingFreeControlRepository) Create(lo *model.TeachingFreeControl) error {
  return nil
}

func (m *MockTeachingFreeControlRepository) FindByID(id string) (*model.TeachingFreeControl, error) {
  return &model.TeachingFreeControl{
		ID:             "1",
		TaskID:         100,
		RobotID:        "robot_001",
		TaskType:       "PickAndPlace",
		VisionSystem:   model.JSON{"camera": "on"},
		ForceControl:   model.JSON{"force_limit": 10},
		AIModel:        model.JSON{"model_name": "control_v1"},
		LearningData:   model.JSON{"dataset": "training_set"},
		SuccessRate:    0.95,
		AdaptationTime: 2.5,
		ErrorRecovery:  model.JSON{"strategy": "retry"},
		PerformanceLog: model.JSON{"log": []int{1, 2, 3}},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
  }, nil
}

func (m *MockTeachingFreeControlRepository) FindAll() ([]model.TeachingFreeControl, error) {
  return []model.TeachingFreeControl{
    {
			ID:             "1",
			TaskID:         100,
			RobotID:        "robot_001",
			TaskType:       "PickAndPlace",
			VisionSystem:   model.JSON{"camera": "on"},
			ForceControl:   model.JSON{"force_limit": 10},
			AIModel:        model.JSON{"model_name": "control_v1"},
			LearningData:   model.JSON{"dataset": "training_set"},
			SuccessRate:    0.95,
			AdaptationTime: 2.5,
			ErrorRecovery:  model.JSON{"strategy": "retry"},
			PerformanceLog: model.JSON{"log": []int{1, 2, 3}},
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
    },
  }, nil
}

func (m *MockTeachingFreeControlRepository) Update(id string, po *model.TeachingFreeControl) error {
    return nil
}

func (m *MockTeachingFreeControlRepository) Delete(id string) error {
    return nil
}

// テスト用ルーター
func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockTeachingFreeControlRepository{}
    mockService := &service.TeachingFreeControlService{Repo: mockRepo}
    ctl := &TeachingFreeControlController{Service: mockService}

    r.POST("/api/teaching_free_control", ctl.AddTeachingFreeControl)
    r.PUT("/api/teaching_free_control/:id", ctl.EditTeachingFreeControl)
    r.DELETE("/api/teaching_free_control/:id", ctl.DeleteTeachingFreeControl)
    return r
}

func TestAddLanguageOptimization(t *testing.T) {
  gin.SetMode(gin.TestMode)
  r := setupRouter()

  body := `{
    "task_id": 100,
    "robot_id": "robot_001",
    "task_type": "PickAndPlace",
    "vision_system": {"camera": "on"},
    "force_control": {"force_limit": 10},
    "ai_model": {"model_name": "control_v1"},
    "learning_data": {"dataset": "training_set"},
    "success_rate": 0.95,
    "adaptation_time": 2.5,
    "error_recovery": {"strategy": "retry"},
    "performance_log": {"log": [1,2,3]}
  }`

  req, _ := http.NewRequest(http.MethodPost, "/api/teaching_free_control", bytes.NewBuffer([]byte(body)))
  req.Header.Set("Content-Type", "application/json")
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Contains(t, w.Body.String(), "Teaching free control added")
}

func TestUpdateTeachingFreeControl(t *testing.T) {
  gin.SetMode(gin.TestMode)
  r := setupRouter()

  body := `{
    "success_rate": 0.97,
    "adaptation_time": 2.2
  }`

  req, _ := http.NewRequest(http.MethodPut, "/api/teaching_free_control/1", bytes.NewBuffer([]byte(body)))
  req.Header.Set("Content-Type", "application/json")
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Contains(t, w.Body.String(), "Teaching free control edited")
}

func TestDeleteTeachingFreeControl(t *testing.T) {
  gin.SetMode(gin.TestMode)
  r := setupRouter()

  req, _ := http.NewRequest(http.MethodDelete, "/api/teaching_free_control/1", nil)
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Contains(t, w.Body.String(), "Teaching free control deleted")
}
