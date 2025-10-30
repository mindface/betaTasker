package model

import (
	"time"
)

// KnowledgeEntity - 分析知識エンティティを繋ぐモデル
type KnowledgeEntity struct {
  ID              string    `gorm:"type:varchar(255);primaryKey" json:"id"`
  TaskID          uint      `json:"task_id" gorm:"index"`
  EntityType      string    `json:"entity_type" gorm:"index"`
  ReferenceID     string    `json:"reference_id" gorm:"index"`
  Domain          string		`json:"domain" gorm:"index"`
  AbstractLevel   string    `json:"abstract_level"`
  Source          string    `json:"source"`
  Tags            JSON      `json:"tags" gorm:"type:jsonb"`
  LinkedEntityIDs JSON      `json:"linked_entity_ids" gorm:"type:jsonb"`
  CreatedAt       time.Time `json:"created_at"`
  UpdatedAt       time.Time `json:"updated_at"`

  Task Task `json:"task" gorm:"foreignKey:TaskID"`
}
