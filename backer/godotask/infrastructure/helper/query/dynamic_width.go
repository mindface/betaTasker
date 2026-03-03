package query

import (
	dtoquery "github.com/godotask/dto/query"
	"gorm.io/gorm"
)

// WithDynamicFzilters
// include + optional ID に応じて WHERE を動的構築
func WithDynamicFilters(q dtoquery.QueryFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, target := range q.Include {
			switch target {
        case dtoquery.FilterUser:
          if q.UserID != nil {
            db = db.Where("user_id = ?", *q.UserID)
          }
        case dtoquery.FilterTask:
          if q.TaskID != nil {
            db = db.Where("task_id = ?", *q.TaskID)
          }
        case dtoquery.FilterMemory:
          if q.MemoryID != nil {
            db = db.Where("memory_id = ?", *q.MemoryID)
          }
			}
		}
		return db
	}
}

func WithDynamicIncludes(includes []dtoquery.FilterTarget) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, inc := range includes {
			switch inc {
			case "pattern":
				db = db.Preload("Patterns")
			case "insight":
				db = db.Preload("Insights")
			case "modeler":
				db = db.Preload("Modelers")
			}
		}
		return db
	}
}
