package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type MemoryService struct {
	Repo *repository.MemoryRepository
}

func (s *MemoryService) CreateMemory(memory *model.Memory) error {
	return s.Repo.Create(memory)
}

func (s *MemoryService) GetMemoryByID(id string) (*model.Memory, error) {
	// 実装例: 文字列IDをintに変換し、リポジトリ経由で取得
	return s.Repo.FindByID(id)
}

func (s *MemoryService) ListMemories() ([]model.Memory, error) {
	return s.Repo.FindAll()
}

func (s *MemoryService) UpdateMemory(id string, memory *model.Memory) error {
	return s.Repo.Update(id, memory)
}

func (s *MemoryService) DeleteMemory(id string) error {
	return s.Repo.Delete(id)
}

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) CreateTask(task *model.Task) error {
	return s.Repo.Create(task)
}
// 他CRUDも同様

type AssessmentService struct {
	Repo *repository.AssessmentRepository
}

func (s *AssessmentService) CreateAssessment(a *model.Assessment) error {
	return s.Repo.Create(a)
}
// 他CRUDも同様
