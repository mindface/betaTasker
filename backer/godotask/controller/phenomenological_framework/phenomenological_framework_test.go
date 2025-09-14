package phenomenological_framework

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
type MockPhenomenologicalFrameworkRepository struct{}

func (m *MockPhenomenologicalFrameworkRepository) Create(po *model.PhenomenologicalFramework) error {
    return nil
}

func (m *MockPhenomenologicalFrameworkRepository) FindByID(id string) (*model.PhenomenologicalFramework, error) {
  return &model.PhenomenologicalFramework{
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

func (m *MockPhenomenologicalFrameworkRepository) FindAll() ([]model.PhenomenologicalFramework, error) {
  return []model.PhenomenologicalFramework{
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

func (m *MockPhenomenologicalFrameworkRepository) Update(id string, po *model.PhenomenologicalFramework) error {
    return nil
}

func (m *MockPhenomenologicalFrameworkRepository) Delete(id string) error {
    return nil
}

// テスト用ルーター
func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockPhenomenologicalFrameworkRepository{}
    mockService := &service.PhenomenologicalFrameworkService{Repo: mockRepo}
    ctl := &PhenomenologicalFrameworkController{Service: mockService}

    r.POST("/api/phenomenologicalframework", ctl.AddPhenomenologicalFramework)
    r.PUT("/api/phenomenologicalframework/:id", ctl.EditPhenomenologicalFramework)
    r.DELETE("/api/phenomenologicalframework/:id", ctl.DeletePhenomenologicalFramework)
    return r
}

func TestAddPhenomenologicalFramework(t *testing.T) {
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

    req, _ := http.NewRequest(http.MethodPost, "/api/phenomenologicalframework", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Process optimization added")
}

func TestUpdatePhenomenologicalFramework(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
        "optimization_type": "accuracy",
        "improvement": 15.0
    }`

    req, _ := http.NewRequest(http.MethodPut, "/api/phenomenologicalframework/1", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Phenomenological framework edited")
}

func TestDeletePhenomenologicalFramework(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/api/phenomenologicalframework/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Phenomenological framework deleted")
}
