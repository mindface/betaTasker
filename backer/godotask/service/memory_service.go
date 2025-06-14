package service

import (
	"fmt"
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
	return s.Repo.FindByID(id)
}

func (s *MemoryService) ListMemories() ([]model.Memory, error) {
  fmt.Println("Listing all memories")
	return s.Repo.FindAll()
}

func (s *MemoryService) UpdateMemory(id string, memory *model.Memory) error {
	return s.Repo.Update(id, memory)
}

func (s *MemoryService) DeleteMemory(id string) error {
	return s.Repo.Delete(id)
}
