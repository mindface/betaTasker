package repository

import (
	"gorm.io/gorm"
  "github.com/godotask/infrastructure/db/model"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"fmt"
)

type TaskRepositoryImpl struct {
	DB *gorm.DB
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

func (r *TaskRepositoryImpl) FindAll(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	// if err := r.DB.Find(&tasks).Error; err != nil {
	// 	return nil, err
	// }

	err := r.DB.
    Scopes(helperquery.WithUserFilter(userID)).
    Preload("Assessments").
    Preload("KnowledgePatterns").
    // Preload("HeuristicsModel").
    // Preload("HeuristicsTracking").
    // Preload("HeuristicsAnalysis").
    // Preload("HeuristicsPattern").
    // Preload("HeuristicsInsight").
    // Preload("LanguageOptimization").
    Order("created_at ASC, id ASC").
    Find(&tasks).Error

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// ListTasksByUser: 特定ユーザーのタスク一覧を取得
func (r *TaskRepositoryImpl) ListTasksPager(filter dtoquery.QueryFilter, offset int, limit int) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	q := r.DB.Model(&model.Task{}).Scopes(helperquery.WithDynamicFilters(filter))

	if q.Error != nil {
		fmt.Printf("Error in Scopes(helperquery.WithDynamicFilters): %v\n", q.Error)
		return nil, 0, q.Error
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.
    Preload("KnowledgePatterns").
		Order("created_at DESC, id DESC").
		Limit(limit).
		Offset(offset).
		Find(&tasks).Error; err != nil {
		return nil, 0, err
	}
  return tasks, total,nil
}

// ListTasksByUserPager: 特定ユーザーのタスク一覧をページネーション取得
func (r *TaskRepositoryImpl) ListTasksByUserPager(userID uint, offset, limit int) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	q := r.DB.Model(&model.Task{}).Scopes(helperquery.WithUserFilter(userID))
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.
		Order("created_at DESC, id DESC").
		Limit(limit).
		Offset(offset).
		Find(&tasks).Error; err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}

func (r *TaskRepositoryImpl) Update(id string, task *model.Task) error {
	return r.DB.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
}

func (r *TaskRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.Task{}, id).Error
}

func (r *TaskRepositoryImpl) ListSearchTasksPager(filter dtoquery.QueryFilter, offset int, limit int) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64
	searchTerm := ""
	if filter.Search != nil {
		searchTerm = *filter.Search
	}
	fmt.Printf("Search Filter Applied - SearchTerm: %s, UserID: %v\n", searchTerm, filter.UserID)

	q := r.DB.Model(&model.Task{}).Scopes(helperquery.WithDynamicFilters(filter))

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.
    Preload("KnowledgePatterns").
		Order("created_at DESC, id DESC").
		Limit(limit).
		Offset(offset).
		Find(&tasks).Error; err != nil {
		return nil, 0, err
	}
  return tasks, total,nil
}
