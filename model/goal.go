package model

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	MatchID   uint           `json:"match_id" gorm:"not null"`
	PlayerID  uint           `json:"player_id" gorm:"not null"`
	TeamID    uint           `json:"team_id" gorm:"not null"`
	Minute    int            `json:"minute" gorm:"not null"`
	Player    *Player        `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
	Team      *Team          `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Match     *Match         `json:"match,omitempty" gorm:"foreignKey:MatchID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
