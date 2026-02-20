package dto

type GoalDetail struct {
	PlayerName string `json:"player_name"`
	TeamName   string `json:"team_name"`
	Minute     int    `json:"minute"`
}

type TopScorer struct {
	PlayerName string `json:"player_name"`
	Goals      int    `json:"goals"`
}

type TeamInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type MatchReport struct {
	MatchID            uint         `json:"match_id"`
	MatchDate          string       `json:"match_date"`
	MatchTime          string       `json:"match_time"`
	HomeTeam           TeamInfo     `json:"home_team"`
	AwayTeam           TeamInfo     `json:"away_team"`
	HomeScore          int          `json:"home_score"`
	AwayScore          int          `json:"away_score"`
	Status             string       `json:"status"`
	Goals              []GoalDetail `json:"goals"`
	TopScorer          *TopScorer   `json:"top_scorer"`
	CumulativeHomeWins int64        `json:"cumulative_home_wins"`
	CumulativeAwayWins int64        `json:"cumulative_away_wins"`
}
