package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/godotask/model"
	"github.com/godotask/service"

	"github.com/stretchr/testify/assert"
)

// --- モックリポジトリの定義 ---
type MockAssessmentRepository struct {
	CreateFunc func(*model.Assessment) error
	FindByIDFunc func(string) (*model.Assessment, error)
	FindAllFunc func() ([]model.Assessment, error)
	UpdateFunc func(string, *model.Assessment) error
	DeleteFunc func(string) error
}

func (m *MockAssessmentRepository) Create(a *model.Assessment) error {
	return m.CreateFunc(a)
}
func (m *MockAssessmentRepository) FindByID(id string) (*model.Assessment, error) {
	return m.FindByIDFunc(id)
}
func (m *MockAssessmentRepository) FindAll() ([]model.Assessment, error) {
	return m.FindAllFunc()
}
func (m *MockAssessmentRepository) Update(id string, a *model.Assessment) error {
	return m.UpdateFunc(id, a)
}
func (m *MockAssessmentRepository) Delete(id string) error {
	return m.DeleteFunc(id)
}


func TestCreateAssessment(t *testing.T) {
	mockRepo := &MockAssessmentRepository{
		CreateFunc: func(a *model.Assessment) error {
			return nil
		},
	}
	svc := service.AssessmentService{Repo: mockRepo}
	assessment := &model.Assessment{ID: 1, UserID: 123}

	err := svc.CreateAssessment(assessment)
	assert.NoError(t, err)
}

func TestGetAssessmentByID(t *testing.T) {
	expected := &model.Assessment{
		ID: 1, UserID: 123, CreatedAt: time.Now(),
	}
	mockRepo := &MockAssessmentRepository{
		FindByIDFunc: func(id string) (*model.Assessment, error) {
			return expected, nil
		},
	}
	svc := service.AssessmentService{Repo: mockRepo}
	result, err := svc.GetAssessmentByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expected.ID, result.ID)
}

func TestListAssessments(t *testing.T) {
	mockRepo := &MockAssessmentRepository{
		FindAllFunc: func() ([]model.Assessment, error) {
			return []model.Assessment{
				{ID: 1}, {ID: 2},
			}, nil
		},
	}
	svc := service.AssessmentService{Repo: mockRepo}
	list, err := svc.ListAssessments()

	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestUpdateAssessment(t *testing.T) {
	mockRepo := &MockAssessmentRepository{
		UpdateFunc: func(id string, a *model.Assessment) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	svc := service.AssessmentService{Repo: mockRepo}

	err := svc.UpdateAssessment("1", &model.Assessment{
		QualitativeFeedback: "Updated feedback",
	})
	assert.NoError(t, err)
}

func TestDeleteAssessment(t *testing.T) {
	mockRepo := &MockAssessmentRepository{
		DeleteFunc: func(id string) error {
			return nil
		},
	}
	svc := service.AssessmentService{Repo: mockRepo}

	err := svc.DeleteAssessment("1")
	assert.NoError(t, err)
}

