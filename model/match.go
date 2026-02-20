package model

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	MatchDate  string         `json:"match_date" gorm:"size:10;not null"`
	MatchTime  string         `json:"match_time" gorm:"size:8;not null"`
	HomeTeamID uint           `json:"home_team_id" gorm:"not null"`
	AwayTeamID uint           `json:"away_team_id" gorm:"not null"`
	HomeScore  *int           `json:"home_score"`
	AwayScore  *int           `json:"away_score"`
	HomeTeam   *Team          `json:"home_team,omitempty" gorm:"foreignKey:HomeTeamID"`
	AwayTeam   *Team          `json:"away_team,omitempty" gorm:"foreignKey:AwayTeamID"`
	Goals      []Goal         `json:"goals,omitempty" gorm:"foreignKey:MatchID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
