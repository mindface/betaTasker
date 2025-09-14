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
      ID:             "1",
      Type:           "tacit",
      Domain:         "AI",
      TacitKnowledge: "Experienced insight",
      ExplicitForm:   "Documented procedure",
      ConversionPath: model.JSON{"from": "tacit", "to": "explicit"},
      Accuracy:       0.95,
      Coverage:       0.9,
      Consistency:    0.92,
      AbstractLevel:  "Medium",
      CreatedAt:      time.Now(),
      UpdatedAt:      time.Now(),
  }, nil
}

func (m *MockKnowledgePatternsRepository) FindAll() ([]model.KnowledgePattern, error) {
  return []model.KnowledgePattern{
    {
			ID:             "1",
			Type:           "tacit",
			Domain:         "AI",
			TacitKnowledge: "Experienced insight",
			ExplicitForm:   "Documented procedure",
			ConversionPath: model.JSON{"from": "tacit", "to": "explicit"},
			Accuracy:       0.95,
			Coverage:       0.9,
			Consistency:    0.92,
			AbstractLevel:  "Medium",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
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
    mockService := &service.KnowledgePatternService{Repo: mockRepo}
    ctl := &KnowledgePatternController{Service: mockService}

    r.POST("/api/knowledge_patterns", ctl.AddKnowledgePattern)
    r.PUT("/api/knowledge_patterns/:id", ctl.EditKnowledgePattern)
    r.DELETE("/api/knowledge_patterns/:id", ctl.DeleteKnowledgePattern)
    return r
}

func TestAddKnowledgePattern(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := setupRouter()

    body := `{
      "type": "tacit",
      "domain": "AI",
      "tacit_knowledge": "Experienced insight",
      "explicit_form": "Documented procedure",
      "conversion_path": {"from":"tacit","to":"explicit"},
      "accuracy": 0.95,
      "coverage": 0.9,
      "consistency": 0.92,
      "abstract_level": "Medium"
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
      "accuracy": 0.97,
      "coverage": 0.92
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
