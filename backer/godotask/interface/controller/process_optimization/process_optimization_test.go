package process_optimization

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/godotask/infrastructure/db/model"
    "github.com/godotask/usecase/service"
)

// モックリポジトリ
type MockProcessOptimizationRepository struct{}

func (m *MockProcessOptimizationRepository) Create(po *model.ProcessOptimization) error {
  return nil
}

func (m *MockProcessOptimizationRepository) FindByID(id string) (*model.ProcessOptimization, error) {
  return &model.ProcessOptimization{
		ID:              "1",
		ProcessID:       "proc_001",
		OptimizationType: "speed",
		InitialState:    model.JSON{"step": 1},
		OptimizedState:  model.JSON{"step": 2},
		Improvement:     10.5,
		Method:          "GradientDescent",
		Iterations:      5,
		ConvergenceTime: 12.3,
		ValidatedBy:     "tester",
		ValidationDate:  time.Now(),
  }, nil
}

func (m *MockProcessOptimizationRepository) FindAll() ([]model.ProcessOptimization, error) {
  return []model.ProcessOptimization{
    {
			ID:              "1",
			ProcessID:       "proc_001",
			OptimizationType: "speed",
			InitialState:    model.JSON{"step": 1},
			OptimizedState:  model.JSON{"step": 2},
			Improvement:     10.5,
			Method:          "GradientDescent",
			Iterations:      5,
			ConvergenceTime: 12.3,
			ValidatedBy:     "tester",
			ValidationDate:  time.Now(),
    },
  }, nil
}

func (m *MockProcessOptimizationRepository) Update(id string, po *model.ProcessOptimization) error {
    return nil
}

func (m *MockProcessOptimizationRepository) Delete(id string) error {
    return nil
}

// テスト用ルーター
func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockProcessOptimizationRepository{}
    mockService := &service.ProcessOptimizationService{Repo: mockRepo}
    ctl := &ProcessOptimizationController{Service: mockService}

    r.POST("/api/processoptimization", ctl.AddProcessOptimization)
    r.PUT("/api/processoptimization/:id", ctl.EditProcessOptimization)
    r.DELETE("/api/processoptimization/:id", ctl.DeleteProcessOptimization)
    return r
}

func TestAddProcessOptimization(t *testing.T) {
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

    req, _ := http.NewRequest(http.MethodPost, "/api/processoptimization", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Process optimization added")
}

func TestUpdateProcessOptimization(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
      "optimization_type": "accuracy",
      "improvement": 15.0
    }`

    req, _ := http.NewRequest(http.MethodPut, "/api/processoptimization/1", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Process optimization edited")
}

func TestDeleteProcessOptimization(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/api/processoptimization/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Process optimization deleted")
}
