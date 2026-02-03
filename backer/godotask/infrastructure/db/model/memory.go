package model

import (
	"time"
)

type Memory struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserID     int       `json:"user_id"`
	SourceType string    `json:"source_type" gorm:"default:book"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Notes      string    `json:"notes"`
	Tags       string    `json:"tags"`
	ReadStatus string    `json:"read_status" gorm:"default:unread"`
	ReadDate   *time.Time `json:"read_date"`
	Factor           string    `json:"factor"`
	Process          string    `json:"process"`
	EvaluationAxis   string    `json:"evaluation_axis"`
	InformationAmount string   `json:"information_amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Memory) TableName() string {
  return "memories"
}