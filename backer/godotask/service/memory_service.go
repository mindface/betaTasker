package service

import (
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/repository"
)

type MemoryService struct {
	Repo repository.MemoryRepositoryInterface
	ContextRepo repository.MemoryContextRepositoryInterface
}

func (s *MemoryService) CreateMemory(memory *model.Memory) error {
	return s.Repo.Create(memory)
}

func (s *MemoryService) GetMemoryByID(id string) (*model.Memory, error) {
	return s.Repo.FindByID(id)
}

func (s *MemoryService) ListMemories() ([]model.Memory, error) {
	return s.Repo.FindAll()
}

func (s *MemoryService) ListMemoriesTOPager(page, perPage, offset int) ([]model.Memory, int64, error) {
    return s.Repo.ListMemories(offset, perPage)
}

func (s *MemoryService) UpdateMemory(id string, memory *model.Memory) error {
	return s.Repo.Update(id, memory)
}

func (s *MemoryService) DeleteMemory(id string) error {
	return s.Repo.Delete(id)
}

func (s *MemoryService) FindMemoryContextsByCode(code string, contexts *[]model.MemoryContext) error {
	return s.ContextRepo.FindByCode(code, contexts)
}

func (s *MemoryService) FindMemoryAidsByCode(code string, contexts *[]model.MemoryContext) error {
	return s.ContextRepo.FindWithAidsByCode(code, contexts)
}
