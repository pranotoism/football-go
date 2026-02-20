package service

import (
	"errors"

	"github.com/pranotoism/football-go/dto"
	"github.com/pranotoism/football-go/model"
	"github.com/pranotoism/football-go/repository"
	"gorm.io/gorm"
)

type ReportService struct {
	matchRepo *repository.MatchRepository
}

func NewReportService(matchRepo *repository.MatchRepository) *ReportService {
	return &ReportService{matchRepo: matchRepo}
}

func (s *ReportService) GetMatchReport(id uint) (*dto.MatchReport, error) {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	if match.HomeScore == nil {
		return nil, errors.New("match result has not been reported yet")
	}

	return s.buildReport(match), nil
}

func (s *ReportService) GetAllMatchReports(page, perPage int) ([]dto.MatchReport, int64, error) {
	matches, total, err := s.matchRepo.FindPlayedMatches(page, perPage)
	if err != nil {
		return nil, 0, err
	}

	reports := make([]dto.MatchReport, len(matches))
	for i := range matches {
		reports[i] = *s.buildReport(&matches[i])
	}

	return reports, total, nil
}

func (s *ReportService) buildReport(match *model.Match) *dto.MatchReport {
	status := "Draw"
	if *match.HomeScore > *match.AwayScore {
		status = "Home Win"
	} else if *match.AwayScore > *match.HomeScore {
		status = "Away Win"
	}

	goals := make([]dto.GoalDetail, len(match.Goals))
	playerGoalCount := make(map[uint]int)
	playerNames := make(map[uint]string)

	for i, g := range match.Goals {
		goals[i] = dto.GoalDetail{
			PlayerName: g.Player.Name,
			TeamName:   g.Team.Name,
			Minute:     g.Minute,
		}
		playerGoalCount[g.PlayerID]++
		playerNames[g.PlayerID] = g.Player.Name
	}

	var topScorer *dto.TopScorer
	maxGoals := 0
	for playerID, count := range playerGoalCount {
		if count > maxGoals {
			maxGoals = count
			topScorer = &dto.TopScorer{
				PlayerName: playerNames[playerID],
				Goals:      count,
			}
		}
	}

	cumulativeHomeWins := s.matchRepo.CountWins(match.HomeTeamID)
	cumulativeAwayWins := s.matchRepo.CountWins(match.AwayTeamID)

	return &dto.MatchReport{
		MatchID:            match.ID,
		MatchDate:          match.MatchDate,
		MatchTime:          match.MatchTime,
		HomeTeam:           dto.TeamInfo{ID: match.HomeTeam.ID, Name: match.HomeTeam.Name},
		AwayTeam:           dto.TeamInfo{ID: match.AwayTeam.ID, Name: match.AwayTeam.Name},
		HomeScore:          *match.HomeScore,
		AwayScore:          *match.AwayScore,
		Status:             status,
		Goals:              goals,
		TopScorer:          topScorer,
		CumulativeHomeWins: cumulativeHomeWins,
		CumulativeAwayWins: cumulativeAwayWins,
	}
}
