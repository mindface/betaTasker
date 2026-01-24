package phenomenological_framework

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
type MockPhenomenologicalFrameworkRepository struct{}

func (m *MockPhenomenologicalFrameworkRepository) Create(po *model.PhenomenologicalFramework) error {
  return nil
}

func (m *MockPhenomenologicalFrameworkRepository) FindByID(id string) (*model.PhenomenologicalFramework, error) {
  return &model.PhenomenologicalFramework{
    ID:            "1",
    TaskID:        100,
    Name:          "Test Framework",
    Description:   "Description of test framework",
    Goal:          "Optimize performance",
    Scope:         "System Level",
    Process:       model.JSON{"step": "input"},
    Result:        model.JSON{"output": "success"},
    Feedback:      model.JSON{"quality": "good"},
    LimitMin:      0.1,
    LimitMax:      99.9,
    GoalFunction:  "f(x)=x^2",
    AbstractLevel: "High",
    Domain:        "AI",
    CreatedAt:     time.Now(),
    UpdatedAt:     time.Now(),
  }, nil
}

func (m *MockPhenomenologicalFrameworkRepository) FindAll() ([]model.PhenomenologicalFramework, error) {
  return []model.PhenomenologicalFramework{
    {
      ID:            "1",
      TaskID:        100,
      Name:          "Test Framework",
      Description:   "Description of test framework",
      Goal:          "Optimize performance",
      Scope:         "System Level",
      Process:       model.JSON{"step": "input"},
      Result:        model.JSON{"output": "success"},
      Feedback:      model.JSON{"quality": "good"},
      LimitMin:      0.1,
      LimitMax:      99.9,
      GoalFunction:  "f(x)=x^2",
      AbstractLevel: "High",
      Domain:        "AI",
      CreatedAt:     time.Now(),
      UpdatedAt:     time.Now(),
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
    "task_id": 100,
    "name": "Framework Alpha",
    "description": "Test description",
    "goal": "Improve accuracy",
    "scope": "System Level",
    "process": {"step": "input"},
    "result": {"output": "success"},
    "feedback": {"note": "stable"},
    "limit_min": 0.5,
    "limit_max": 99.5,
    "goal_function": "f(x)=x^2",
    "abstract_level": "Medium",
    "domain": "AI"
  }`

  req, _ := http.NewRequest(http.MethodPost, "/api/phenomenologicalframework", bytes.NewBuffer([]byte(body)))
  req.Header.Set("Content-Type", "application/json")
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Contains(t, w.Body.String(), "Phenomenological framework added")
}

func TestUpdatePhenomenologicalFramework(t *testing.T) {
  gin.SetMode(gin.TestMode)
  r := setupRouter()

  body := `{
    "description": "Updated description",
    "goal": "Maximize throughput"
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
