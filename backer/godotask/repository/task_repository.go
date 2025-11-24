package repository

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func attachEntity(db *gorm.DB, ke *model.KnowledgeEntity) error {
    switch ke.EntityType {
			case "heuristics_analysis":
					// var ha model.HeuristicsAnalysis
					// if err := db.Where("id = ?", ke.ReferenceID).First(&ha).Error; err != nil {
					// 		return err
					// }
					var ha model.HeuristicsAnalysis
					err := db.Unscoped().
							Where("id = ?", ke.ReferenceID).
							First(&ha).Error

					if err != nil {
							// 削除済みでも無視したい場合はログのみにする
							fmt.Printf("[WARN] HeuristicsAnalysis not found (may be soft-deleted). id=%s", ke.ReferenceID)
							ke.Entity = nil
							return nil
					}
					ke.Entity = ha
			case "heuristics_insight":
					var hi model.HeuristicsInsight
					err := db.Unscoped().
							Where("id = ?", ke.ReferenceID).
							First(&hi).Error

					if err != nil {
							// 削除済みでも無視したい場合はログのみにする
							fmt.Printf("[WARN] HeuristicsInsight not found (may be soft-deleted). id=%s", ke.ReferenceID)
							ke.Entity = nil
							return nil
					}
					ke.Entity = hi
			case "memory_context":
					var mc model.MemoryContext
					if err := db.Where("id = ?", ke.ReferenceID).First(&mc).Error; err != nil {
							return err
					}
					ke.Entity = mc
			case "optimization_model":
					var om model.OptimizationModel
					if err := db.Where("id = ?", ke.ReferenceID).First(&om).Error; err != nil {
							return err
					}
					ke.Entity = om
			case "phenomenological_framework":
					var pf model.PhenomenologicalFramework
					err := db.Unscoped().
							Where("id = ?", ke.ReferenceID).
							First(&pf).Error

					if err != nil {
							// 削除済みでも無視したい場合はログのみにする
							fmt.Printf("[WARN] PhenomenologicalFramework not found (may be soft-deleted). id=%s", ke.ReferenceID)
							ke.Entity = nil
							return nil
					}
					ke.Entity = pf
			case "quantification_label":
					var ql model.QuantificationLabel
					err := db.Unscoped().
							Where("id = ?", ke.ReferenceID).
							First(&ql).Error

					if err != nil {
							// 削除済みでも無視したい場合はログのみにする
							fmt.Printf("[WARN] QuantificationLabel not found (may be soft-deleted). id=%s", ke.ReferenceID)
							ke.Entity = nil
							return nil
					}
					ke.Entity = ql
			default:
					// 未対応の entity_type はログに出して無視する
					fmt.Printf("[WARN] Unknown entity_type: %s for KnowledgeEntity ID=%s", ke.EntityType, ke.ID)
    }
    return nil
}

func (r *TaskRepositoryImpl) Create(task *model.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepositoryImpl) FindByID(id string) (*model.Task, error) {
	var task model.Task
	if err := r.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	err := r.DB.Preload("QualitativeLabels").
		// Preload("QuantificationLabels").
		// Preload("MultimodalData").
		Preload("KnowledgeEntities").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
    // KnowledgeEntity の entity_type ごとに中身を埋める
	for ti := range tasks {
		for ki := range tasks[ti].KnowledgeEntities {
			err := attachEntity(r.DB, &tasks[ti].KnowledgeEntities[ki])
			if err != nil {
				return nil, err
			}
		}
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) Update(id string, task *model.Task) error {
	return r.DB.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
}

func (r *TaskRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.Task{}, id).Error
}
