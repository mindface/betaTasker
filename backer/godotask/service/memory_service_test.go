package service_test

import (
	"testing"

	"github.com/godotask/model"
	"github.com/godotask/service"
	"github.com/stretchr/testify/assert"
)

// --- モック MemoryRepository ---
type MockMemoryRepository struct {
	CreateFunc func(*model.Memory) error
	FindByIDFunc func(string) (*model.Memory, error)
	FindAllFunc func() ([]model.Memory, error)
	UpdateFunc func(string, *model.Memory) error
	DeleteFunc func(string) error
}

func (m *MockMemoryRepository) Create(memory *model.Memory) error {
	return m.CreateFunc(memory)
}
func (m *MockMemoryRepository) FindByID(id string) (*model.Memory, error) {
	return m.FindByIDFunc(id)
}
func (m *MockMemoryRepository) FindAll() ([]model.Memory, error) {
	return m.FindAllFunc()
}
func (m *MockMemoryRepository) Update(id string, memory *model.Memory) error {
	return m.UpdateFunc(id, memory)
}
func (m *MockMemoryRepository) Delete(id string) error {
	return m.DeleteFunc(id)
}

// --- モック MemoryContextRepository ---
type MockMemoryContextRepository struct {
	FindByCodeFunc func(string, *[]model.MemoryContext) error
	FindWithAidsByCodeFunc func(string, *[]model.MemoryContext) error
}

func (m *MockMemoryContextRepository) FindByCode(code string, contexts *[]model.MemoryContext) error {
	return m.FindByCodeFunc(code, contexts)
}
func (m *MockMemoryContextRepository) FindWithAidsByCode(code string, contexts *[]model.MemoryContext) error {
	return m.FindWithAidsByCodeFunc(code, contexts)
}

// --- テスト関数 ---

func TestCreateMemory(t *testing.T) {
	mockRepo := &MockMemoryRepository{
		CreateFunc: func(m *model.Memory) error {
			return nil
		},
	}
	service := service.MemoryService{Repo: mockRepo}

	err := service.CreateMemory(&model.Memory{ID: 1})
	assert.NoError(t, err)
}

func TestGetMemoryByID(t *testing.T) {
	expected := &model.Memory{ID: 1}
	mockRepo := &MockMemoryRepository{
		FindByIDFunc: func(id string) (*model.Memory, error) {
			return expected, nil
		},
	}
	service := service.MemoryService{Repo: mockRepo}

	result, err := service.GetMemoryByID("1")
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, result.ID)
}

func TestListMemories(t *testing.T) {
	mockRepo := &MockMemoryRepository{
		FindAllFunc: func() ([]model.Memory, error) {
			return []model.Memory{{ID: 1}, {ID: 2}}, nil
		},
	}
	service := service.MemoryService{Repo: mockRepo}

	list, err := service.ListMemories()
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestUpdateMemory(t *testing.T) {
	mockRepo := &MockMemoryRepository{
		UpdateFunc: func(id string, m *model.Memory) error {
			return nil
		},
	}
	service := service.MemoryService{Repo: mockRepo}

	err := service.UpdateMemory("1", &model.Memory{})
	assert.NoError(t, err)
}

func TestDeleteMemory(t *testing.T) {
	mockRepo := &MockMemoryRepository{
		DeleteFunc: func(id string) error {
			return nil
		},
	}
	service := service.MemoryService{Repo: mockRepo}

	err := service.DeleteMemory("1")
	assert.NoError(t, err)
}

func TestFindMemoryContextsByCode(t *testing.T) {
	mockContextRepo := &MockMemoryContextRepository{
		FindByCodeFunc: func(code string, ctxs *[]model.MemoryContext) error {
			*ctxs = []model.MemoryContext{{ID: 1}, {ID: 2}}
			return nil
		},
	}
	service := service.MemoryService{ContextRepo: mockContextRepo}

	var result []model.MemoryContext
	err := service.FindMemoryContextsByCode("code123", &result)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestFindMemoryAidsByCode(t *testing.T) {
	mockContextRepo := &MockMemoryContextRepository{
		FindWithAidsByCodeFunc: func(code string, ctxs *[]model.MemoryContext) error {
			*ctxs = []model.MemoryContext{{ID: 1}}
			return nil
		},
	}
	service := service.MemoryService{ContextRepo: mockContextRepo}

	var result []model.MemoryContext
	err := service.FindMemoryAidsByCode("code123", &result)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}
