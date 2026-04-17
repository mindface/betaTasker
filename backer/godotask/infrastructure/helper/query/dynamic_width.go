package query

import (
	dtoquery "github.com/godotask/dto/query"
	"gorm.io/gorm"
)

// Include + optional ID + Search に応じて WHERE を動的構築
func WithDynamicFilters(q dtoquery.QueryFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q.UserID != nil && *q.UserID != 0 {
			db = db.Where("user_id = ?", *q.UserID)
		}

		if q.TaskID != nil {
			db = db.Where("task_id = ?", *q.TaskID)
		}

		if q.MemoryID != nil {
			db = db.Where("memory_id = ?", *q.MemoryID)
		}

		if q.Search != nil && *q.Search != "" {
			keyword := "%" + *q.Search + "%"
			db = db.Where("title LIKE ? OR description LIKE ?", keyword, keyword)
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
