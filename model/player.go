package model

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	TeamID       uint           `json:"team_id" gorm:"not null"`
	Name         string         `json:"name" gorm:"size:255;not null"`
	HeightCM     int            `json:"height_cm"`
	WeightKG     int            `json:"weight_kg"`
	Position     string         `json:"position" gorm:"type:varchar(20);not null"`
	JerseyNumber int            `json:"jersey_number" gorm:"not null"`
	Team         *Team          `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
