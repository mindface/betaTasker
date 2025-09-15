package model

import (
	"time"
)

type User struct {
	ID                uint      `gorm:"primaryKey"`
	Username          string    `gorm:"unique;not null"`
	Email             string    `gorm:"unique;not null"`
	PasswordHash      string    `gorm:"not null"`
	Role              string    `gorm:"default:'user'"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	IsActive          bool      `gorm:"default:true"`
	Factor            string    `json:"factor"`
	Process           string    `json:"process"`
	EvaluationAxis    string    `json:"evaluation_axis"`
	InformationAmount string    `json:"information_amount"`

	Tasks               []Task               `json:"tasks" gorm:"foreignKey:UserID"`
	QuantificationLabels []QuantificationLabel `json:"labels" gorm:"foreignKey:UserID"`
	MultimodalData      []MultimodalData     `json:"multimodal_data" gorm:"foreignKey:UserID"`
}