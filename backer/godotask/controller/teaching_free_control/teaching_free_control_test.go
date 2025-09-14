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
    ID:               "1",
    ProcessID:        "proc_001",
    OptimizationType: "speed",
    InitialState:     model.JSON{"step": 1},
    OptimizedState:   model.JSON{"step": 2},
    Improvement:      10.5,
    Method:           "GradientDescent",
    Iterations:       5,
    ConvergenceTime:  12.3,
    ValidatedBy:      "tester",
    ValidationDate:   time.Now(),
  }, nil
}

func (m *MockTeachingFreeControlRepository) FindAll() ([]model.TeachingFreeControl, error) {
  return []model.TeachingFreeControl{
    {
      ID:               "1",
      ProcessID:        "proc_001",
      OptimizationType: "speed",
      InitialState:     model.JSON{"step": 1},
      OptimizedState:   model.JSON{"step": 2},
      Improvement:      10.5,
      Method:           "GradientDescent",
      Iterations:       5,
      ConvergenceTime:  12.3,
      ValidatedBy:      "tester",
      ValidationDate:   time.Now(),
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
        "process_id": "proc_001",
        "optimization_type": "speed",
        "initial_state": {"step":1},
        "optimized_state": {"step":2},
        "improvement": 10.5,
        "method": "GradientDescent",
        "iterations": 5,
        "convergence_time": 12.3,
        "validated_by": "tester",
        "validation_date": "2025-09-13T12:00:00Z"
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
        "optimization_type": "accuracy",
        "improvement": 15.0
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
