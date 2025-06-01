package model

import (
	"time"
)

type Memory struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserID     int       `json:"user_id"`
	SourceType string    `json:"source_type"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Notes      string    `json:"notes"`
	Tags       string    `json:"tags"`
	ReadStatus string    `json:"read_status"`
	ReadDate   *time.Time `json:"read_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Memory) TableName() string {
    return "memory"
}