package knowledge_pattern

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
type MockKnowledgePatternsRepository struct{}

func (m *MockKnowledgePatternsRepository) Create(knowledgePattern *model.KnowledgePattern) error {
  return nil
}

func (m *MockKnowledgePatternsRepository) FindByID(id string) (*model.KnowledgePattern, error) {
  return &model.KnowledgePattern{
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

func (m *MockKnowledgePatternsRepository) FindAll() ([]model.KnowledgePattern, error) {
  return []model.KnowledgePattern{
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

func (m *MockKnowledgePatternsRepository) Update(id string, knowledgePattern *model.KnowledgePattern) error {
    return nil
}

func (m *MockKnowledgePatternsRepository) Delete(id string) error {
    return nil
}

// テスト用ルーター
func setupRouter() *gin.Engine {
    r := gin.Default()
    mockRepo := &MockKnowledgePatternsRepository{}
    mockService := &service.KnowledgePatternsService{Repo: mockRepo}
    ctl := &KnowledgePatternsController{Service: mockService}

    r.POST("/api/knowledge_patterns", ctl.AddKnowledgePattern)
    r.PUT("/api/knowledge_patterns/:id", ctl.EditKnowledgePattern)
    r.DELETE("/api/knowledge_patterns/:id", ctl.DeleteKnowledgePattern)
    return r
}

func TestAddKnowledgePattern(t *testing.T) {
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

    req, _ := http.NewRequest(http.MethodPost, "/api/knowledge_patterns", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Knowledge pattern added")
}

func TestUpdateKnowledgePattern(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
        "optimization_type": "accuracy",
        "improvement": 15.0
    }`

    req, _ := http.NewRequest(http.MethodPut, "/api/knowledge_patterns/1", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Knowledge pattern edited")
}

func TestDeleteKnowledgePattern(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/api/knowledge_patterns/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Knowledge pattern deleted")
}
