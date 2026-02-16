package query

import (
	querydto "github.com/godotask/dto/query"
	"gorm.io/gorm"
)

// WithDynamicFilters
// include + optional ID に応じて WHERE を動的構築
func WithDynamicFilters(q querydto.QueryFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, target := range q.Include {
			switch target {
        case querydto.FilterUser:
          if q.UserID != nil {
            db = db.Where("user_id = ?", *q.UserID)
          }
        case querydto.FilterTask:
          if q.TaskID != nil {
            db = db.Where("task_id = ?", *q.TaskID)
          }
        case querydto.FilterMemory:
          if q.MemoryID != nil {
            db = db.Where("memory_id = ?", *q.MemoryID)
          }
			}
		}
		return db
	}
}
