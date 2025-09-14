package qualitative_label

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
type MockQualitativeLabelRepository struct{}

func (m *MockQualitativeLabelRepository) Create(ql *model.QualitativeLabel) error {
  return nil
}

func (m *MockQualitativeLabelRepository) FindByID(id string) (*model.QualitativeLabel, error) {
  return &model.QualitativeLabel{
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

func (m *MockQualitativeLabelRepository) FindAll() ([]model.QualitativeLabel, error) {
  return []model.QualitativeLabel{
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

func (m *MockQualitativeLabelRepository) Update(id string, ql *model.QualitativeLabel) error {
    return nil
}

func (m *MockQualitativeLabelRepository) Delete(id string) error {
    return nil
}

// テスト用ルーター
func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockQualitativeLabelRepository{}
    mockService := &service.QualitativeLabelService{Repo: mockRepo}
    ctl := &QualitativeLabelController{Service: mockService}

    r.POST("/api/qualitative_label", ctl.AddQualitativeLabel)
    r.PUT("/api/qualitative_label/:id", ctl.EditQualitativeLabel)
    r.DELETE("/api/qualitative_label/:id", ctl.DeleteQualitativeLabel)
    return r
}

func TestAddQualitativeLabel(t *testing.T) {
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

    req, _ := http.NewRequest(http.MethodPost, "/api/qualitative_label", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Qualitative label added")
}

func TestUpdateQualitativeLabel(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
        "optimization_type": "accuracy",
        "improvement": 15.0
    }`

    req, _ := http.NewRequest(http.MethodPut, "/api/qualitative_label/1", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Process optimization edited")
}

func TestDeleteQualitativeLabel(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/api/qualitative_label/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Qualitative label deleted")
}
