package model

import (
	"time"
)

// QualitativeLabel - プロセス最適化記録
type QualitativeLabel struct {
  ID int `gorm:"type:varchar(255);primaryKey" json:"id"`
  TaskID    int            `gorm:"index"`
  UserID    int           `gorm:"index"`
  Content   string         `gorm:"type:text"` // ラベル内容
  Category  string         `gorm:"index"`     // ラベルのカテゴリ
  CreatedAt time.Time
  UpdatedAt time.Time
}
