package model

import "time"

type Task struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	UserID      int        `json:"user_id"`
	MemoryID    *int       `json:"memory_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date"`
	Status      string     `json:"status"` // todo, in_progress, completed
	Priority    int        `json:"priority"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (Task) TableName() string {
    return "task"
}