
package repository

import "github.com/godotask/model"

// BookRepository インターフェース
type BookRepositoryInterface interface {
	Create(book *model.Book) error
	FindByID(id string) (*model.Book, error)
	FindAll() ([]model.Book, error)
	Update(id string, book *model.Book) error
	Delete(id string) error
}

// AssessmentRepository インターフェース
type AssessmentRepositoryInterface interface {
	Create(assessment *model.Assessment) error
	FindByID(id string) (*model.Assessment, error)
	FindAll() ([]model.Assessment, error)
	Update(id string, assessment *model.Assessment) error
	Delete(id string) error
}

type MemoryRepositoryInterface interface {
	Create(memory *model.Memory) error
	FindByID(id string) (*model.Memory, error)
	FindAll() ([]model.Memory, error)
	Update(id string, memory *model.Memory) error
	Delete(id string) error
}

type MemoryContextRepositoryInterface interface {
	FindByCode(code string, contexts *[]model.MemoryContext) error
	FindWithAidsByCode(code string, contexts *[]model.MemoryContext) error
}

type TaskRepositoryInterface interface {
	Create(task *model.Task) error
	FindByID(id string) (*model.Task, error)
	FindAll() ([]model.Task, error)
	Update(id string, task *model.Task) error
	Delete(id string) error
}

