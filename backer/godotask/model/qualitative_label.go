package model

import (
	"time"
)

// QualitativeLabel - プロセス最適化記録
type QualitativeLabel struct {
  ID        string         `gorm:"primaryKey"`
  TaskID    int            `gorm:"index"`
  UserID    uint           `gorm:"index"`
  Content   string         `gorm:"type:text"` // ラベル内容
  Category  string         `gorm:"index"`     // ラベルのカテゴリ
  CreatedAt time.Time
  UpdatedAt time.Time
}
