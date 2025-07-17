package service

import (
	"fmt"
	"github.com/godotask/model"
	"github.com/godotask/repository"
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
  fmt.Println("Listing all memories")
	return s.Repo.FindAll()
}

func (s *MemoryService) UpdateMemory(id string, memory *model.Memory) error {
	return s.Repo.Update(id, memory)
}

func (s *MemoryService) DeleteMemory(id string) error {
	return s.Repo.Delete(id)
}

// // FindMemoryContextsByCode: work_targetにcodeが含まれるものを配列で返す
// func (s *MemoryService) FindMemoryContextsByCode(code string, contexts *[]model.MemoryContext) error {
// 	repo := repository.MemoryContextRepository{DB: s.Repo.DB}
// 	return repo.FindByCode(code, contexts)
// }

// // FindMemoryAidsByCode: work_targetにcodeが含まれるものを補助情報ごと返す
// func (s *MemoryService) FindMemoryAidsByCode(code string, contexts *[]model.MemoryContext) error {
// 	repo := repository.MemoryContextRepository{DB: s.Repo.DB}
// 	return repo.FindWithAidsByCode(code, contexts)
// }


func (s *MemoryService) FindMemoryContextsByCode(code string, contexts *[]model.MemoryContext) error {
	return s.ContextRepo.FindByCode(code, contexts)
}

func (s *MemoryService) FindMemoryAidsByCode(code string, contexts *[]model.MemoryContext) error {
	return s.ContextRepo.FindWithAidsByCode(code, contexts)
}