package dto

type CreateMatchRequest struct {
	MatchDate  string `json:"match_date" binding:"required"`
	MatchTime  string `json:"match_time" binding:"required"`
	HomeTeamID uint   `json:"home_team_id" binding:"required"`
	AwayTeamID uint   `json:"away_team_id" binding:"required"`
}

type UpdateMatchRequest struct {
	MatchDate  string `json:"match_date"`
	MatchTime  string `json:"match_time"`
	HomeTeamID uint   `json:"home_team_id"`
	AwayTeamID uint   `json:"away_team_id"`
}

type GoalInput struct {
	PlayerID uint `json:"player_id" binding:"required"`
	TeamID   uint `json:"team_id" binding:"required"`
	Minute   int  `json:"minute" binding:"required,min=1"`
}

type ReportResultRequest struct {
	HomeScore int         `json:"home_score" binding:"min=0"`
	AwayScore int         `json:"away_score" binding:"min=0"`
	Goals     []GoalInput `json:"goals"`
}
