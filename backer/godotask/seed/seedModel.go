package seed

import (
	// "fmt"
	// "log"
	// "math/rand"
	"time"
	// // "strings" 

	// "gorm.io/gorm"
)

type Book struct {
	ID      int    `gorm:"primaryKey"`
	Title   string
	Name    string
	Text    string
	Disc    string
	ImgPath string
	Status  string
}

type Memory struct {
	ID                 int       `gorm:"primaryKey"`
	UserID             int
	SourceType         string
	Title              string
	Author             string
	Notes              string
	Tags               string
	ReadStatus         string
	ReadDate           *time.Time
	Factor             string
	Process            string
	EvaluationAxis     string
	InformationAmount  string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Task struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int
	MemoryID    int
	Title       string
	Description string
	Date        time.Time
	Status      string
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Assessment struct {
	ID                  int       `gorm:"primaryKey"`
	TaskID              int
	UserID              int
	EffectivenessScore  int
	EffortScore         int
	ImpactScore         int
	QualitativeFeedback string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
