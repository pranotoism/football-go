package model

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:255;not null"`
	LogoURL     string         `json:"logo_url" gorm:"size:500"`
	FoundedYear int            `json:"founded_year" gorm:"not null"`
	HQAddress   string         `json:"hq_address" gorm:"type:text"`
	HQCity      string         `json:"hq_city" gorm:"size:255"`
	Players     []Player       `json:"players,omitempty" gorm:"foreignKey:TeamID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
