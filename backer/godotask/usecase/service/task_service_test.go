package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/usecase/service"
)

// --- モックリポジトリ定義 ---
type MockTaskRepository struct {
	CreateFunc   func(task *model.Task) error
	FindByIDFunc func(id string) (*model.Task, error)
	FindAllFunc  func() ([]model.Task, error)
	UpdateFunc   func(id string, task *model.Task) error
	DeleteFunc   func(id string) error
}

func (m *MockTaskRepository) Create(task *model.Task) error {
	return m.CreateFunc(task)
}
func (m *MockTaskRepository) FindByID(id string) (*model.Task, error) {
	return m.FindByIDFunc(id)
}
func (m *MockTaskRepository) FindAll() ([]model.Task, error) {
	return m.FindAllFunc()
}
func (m *MockTaskRepository) Update(id string, task *model.Task) error {
	return m.UpdateFunc(id, task)
}
func (m *MockTaskRepository) Delete(id string) error {
	return m.DeleteFunc(id)
}

func TestCreateTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		CreateFunc: func(task *model.Task) error {
			return nil
		},
	}
	svc := service.TaskService{Repo: mockRepo}
	task := &model.Task{Title: "Test Task", UserID: 1}

	err := svc.CreateTask(task)
	assert.NoError(t, err)
}

func TestGetTaskByID(t *testing.T) {
	expected := &model.Task{
		ID:     1,
		UserID: 123,
		Title:  "Read book",
	}
	mockRepo := &MockTaskRepository{
		FindByIDFunc: func(id string) (*model.Task, error) {
			return expected, nil
		},
	}
	svc := service.TaskService{Repo: mockRepo}

	result, err := svc.GetTaskByID("1")
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, result.Title)
}

func TestListTasks(t *testing.T) {
	mockRepo := &MockTaskRepository{
		FindAllFunc: func() ([]model.Task, error) {
			return []model.Task{
				{ID: 1, Title: "Task 1"},
				{ID: 2, Title: "Task 2"},
			}, nil
		},
	}
	svc := service.TaskService{Repo: mockRepo}

	tasks, err := svc.ListTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		UpdateFunc: func(id string, task *model.Task) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	svc := service.TaskService{Repo: mockRepo}

	err := svc.UpdateTask("1", &model.Task{Title: "Updated Title"})
	assert.NoError(t, err)

	err = svc.UpdateTask("99", &model.Task{Title: "Non-existent"})
	assert.Error(t, err)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		DeleteFunc: func(id string) error {
			return nil
		},
	}
	svc := service.TaskService{Repo: mockRepo}

	err := svc.DeleteTask("1")
	assert.NoError(t, err)
}
