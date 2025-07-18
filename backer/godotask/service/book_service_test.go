package service_test

import (
	"errors"
	"testing"

	"github.com/godotask/model"
	"github.com/godotask/service"
	"github.com/stretchr/testify/assert"
)

// --- モックリポジトリ定義 ---
type MockBookRepository struct {
	CreateFunc   func(*model.Book) error
	FindByIDFunc func(string) (*model.Book, error)
	FindAllFunc  func() ([]model.Book, error)
	UpdateFunc   func(string, *model.Book) error
	DeleteFunc   func(string) error
}

func (m *MockBookRepository) Create(book *model.Book) error {
	return m.CreateFunc(book)
}
func (m *MockBookRepository) FindByID(id string) (*model.Book, error) {
	return m.FindByIDFunc(id)
}
func (m *MockBookRepository) FindAll() ([]model.Book, error) {
	return m.FindAllFunc()
}
func (m *MockBookRepository) Update(id string, book *model.Book) error {
	return m.UpdateFunc(id, book)
}
func (m *MockBookRepository) Delete(id string) error {
	return m.DeleteFunc(id)
}


func TestCreateBook(t *testing.T) {
	mockRepo := &MockBookRepository{
		CreateFunc: func(b *model.Book) error {
			return nil
		},
	}
	svc := service.BookService{Repo: mockRepo}

	err := svc.CreateBook(&model.Book{Title: "Test Book"})
	assert.NoError(t, err)
}

func TestGetBookByID(t *testing.T) {
	expected := &model.Book{ID: 1, Title: "Book Title"}
	mockRepo := &MockBookRepository{
		FindByIDFunc: func(id string) (*model.Book, error) {
			return expected, nil
		},
	}
	svc := service.BookService{Repo: mockRepo}

	result, err := svc.GetBookByID("1")
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, result.Title)
}

func TestListBooks(t *testing.T) {
	mockRepo := &MockBookRepository{
		FindAllFunc: func() ([]model.Book, error) {
			return []model.Book{
				{ID: 1, Title: "Book A"},
				{ID: 2, Title: "Book B"},
			}, nil
		},
	}
	svc := service.BookService{Repo: mockRepo}

	books, err := svc.ListBooks()
	assert.NoError(t, err)
	assert.Len(t, books, 2)
}

func TestUpdateBook(t *testing.T) {
	mockRepo := &MockBookRepository{
		UpdateFunc: func(id string, b *model.Book) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	svc := service.BookService{Repo: mockRepo}

	err := svc.UpdateBook("1", &model.Book{Name: "Updated Name"})
	assert.NoError(t, err)
}

func TestDeleteBook(t *testing.T) {
	mockRepo := &MockBookRepository{
		DeleteFunc: func(id string) error {
			return nil
		},
	}
	svc := service.BookService{Repo: mockRepo}

	err := svc.DeleteBook("1")
	assert.NoError(t, err)
}
