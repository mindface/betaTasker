package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type KnowledgeEntityService struct {
	Repo repository.KnowledgeEntityRepositoryInterface
	KnowledgeEntityService *KnowledgeEntityService
}

// CreateKnowledgeEntity adds a new knowledge entity
func (s *KnowledgeEntityService) CreateKnowledgeEntity(entity *model.KnowledgeEntity) error {
	return s.Repo.Create(entity)
}

// GetKnowledgeEntityByID retrieves a single entity by ID
func (s *KnowledgeEntityService) GetKnowledgeEntityByID(id string) (*model.KnowledgeEntity, error) {
	return s.Repo.FindByID(id)
}

// ListKnowledgeEntities retrieves all entities
func (s *KnowledgeEntityService) ListKnowledgeEntities() ([]model.KnowledgeEntity, error) {
	return s.Repo.FindAll()
}

// GetKnowledgeEntitiesByTask retrieves entities linked to a specific Task
func (s *KnowledgeEntityService) GetKnowledgeEntitiesByTask(taskID uint) ([]model.KnowledgeEntity, error) {
	return s.Repo.FindByTaskID(taskID)
}

// UpdateKnowledgeEntity updates an entity by ID
func (s *KnowledgeEntityService) UpdateKnowledgeEntity(id string, entity *model.KnowledgeEntity) error {
	return s.Repo.Update(id, entity)
}

// DeleteKnowledgeEntity deletes an entity
func (s *KnowledgeEntityService) DeleteKnowledgeEntity(id string) error {
	return s.Repo.Delete(id)
}

// LinkKnowledgeEntities links two entities bidirectionally
func (s *KnowledgeEntityService) LinkKnowledgeEntities(sourceID, targetID string) error {
	return s.Repo.LinkEntities(sourceID, targetID)
}

// GetLinkedEntities retrieves all entities linked to the specified one
func (s *KnowledgeEntityService) GetLinkedEntities(entityID string) ([]model.KnowledgeEntity, error) {
	return s.Repo.GetLinkedEntities(entityID)
}
