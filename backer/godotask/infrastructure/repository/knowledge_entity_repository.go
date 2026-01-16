package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
	"encoding/json"
)

// KnowledgeEntityRepositoryInterface defines repository methods for KnowledgeEntity
type KnowledgeEntityRepositoryInterface interface {
	Create(entity *model.KnowledgeEntity) error
	FindByID(id string) (*model.KnowledgeEntity, error)
	FindByTaskID(taskID uint) ([]model.KnowledgeEntity, error)
	FindAll() ([]model.KnowledgeEntity, error)
	Update(id string, entity *model.KnowledgeEntity) error
	Delete(id string) error
	LinkEntities(sourceID, targetID string) error
	GetLinkedEntities(entityID string) ([]model.KnowledgeEntity, error)
}

// KnowledgeEntityRepositoryImpl is the concrete implementation
type KnowledgeEntityRepositoryImpl struct {
	DB *gorm.DB
}

// Create inserts a new KnowledgeEntity
func (r *KnowledgeEntityRepositoryImpl) Create(entity *model.KnowledgeEntity) error {
	return r.DB.Create(entity).Error
}

// FindByID retrieves a KnowledgeEntity by its ID
func (r *KnowledgeEntityRepositoryImpl) FindByID(id string) (*model.KnowledgeEntity, error) {
	var entity model.KnowledgeEntity
	if err := r.DB.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindByTaskID returns all KnowledgeEntities related to a Task
func (r *KnowledgeEntityRepositoryImpl) FindByTaskID(taskID uint) ([]model.KnowledgeEntity, error) {
	var entities []model.KnowledgeEntity
	if err := r.DB.Where("task_id = ?", taskID).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// FindAll retrieves all KnowledgeEntities
func (r *KnowledgeEntityRepositoryImpl) FindAll() ([]model.KnowledgeEntity, error) {
	var entities []model.KnowledgeEntity
	if err := r.DB.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Update modifies an existing KnowledgeEntity
func (r *KnowledgeEntityRepositoryImpl) Update(id string, entity *model.KnowledgeEntity) error {
	return r.DB.Model(&model.KnowledgeEntity{}).Where("id = ?", id).Updates(entity).Error
}

// Delete removes a KnowledgeEntity by ID
func (r *KnowledgeEntityRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.KnowledgeEntity{}, id).Error
}

// LinkEntities links two KnowledgeEntities by appending their IDs in LinkedEntityIDs JSON field
func (r *KnowledgeEntityRepositoryImpl) LinkEntities(sourceID, targetID string) error {
	var source, target model.KnowledgeEntity

	if err := r.DB.First(&source, "id = ?", sourceID).Error; err != nil {
		return err
	}
	if err := r.DB.First(&target, "id = ?", targetID).Error; err != nil {
		return err
	}

	// JSON配列を扱う
	// var sourceLinks, targetLinks []string
	// _ = source.LinkedEntityIDs.Unmarshal(&sourceLinks)
	// _ = target.LinkedEntityIDs.Unmarshal(&targetLinks)

	// if !contains(sourceLinks, targetID) {
	// 	sourceLinks = append(sourceLinks, targetID)
	// }
	// if !contains(targetLinks, sourceID) {
	// 	targetLinks = append(targetLinks, sourceID)
	// }

	// source.LinkedEntityIDs.Marshal(sourceLinks)
	// target.LinkedEntityIDs.Marshal(targetLinks)

	if err := r.DB.Save(&source).Error; err != nil {
		return err
	}
	if err := r.DB.Save(&target).Error; err != nil {
		return err
	}
	return nil
}

// GetLinkedEntities returns all KnowledgeEntities linked to a given entity
func (r *KnowledgeEntityRepositoryImpl) GetLinkedEntities(entityID string) ([]model.KnowledgeEntity, error) {
	var entity model.KnowledgeEntity
	if err := r.DB.First(&entity, "id = ?", entityID).Error; err != nil {
		return nil, err
	}

	var linkedIDs []string
if entity.LinkedEntityIDs != nil {
    // JSON を []byte に変換
    b, err := json.Marshal(entity.LinkedEntityIDs)
    if err != nil {
        return nil, err
    }
    // []byte を Go のスライスに変換
    if err := json.Unmarshal(b, &linkedIDs); err != nil {
        return nil, err
    }
}

	var linkedEntities []model.KnowledgeEntity
	if len(linkedIDs) > 0 {
		if err := r.DB.Where("id IN ?", linkedIDs).Find(&linkedEntities).Error; err != nil {
			return nil, err
		}
	}

	return linkedEntities, nil
}

// Utility function
func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

// Constructor
func NewKnowledgeEntityRepository(db *gorm.DB) KnowledgeEntityRepositoryInterface {
	return &KnowledgeEntityRepositoryImpl{DB: db}
}
