package service

import (
	"errors"
	"fmt"

	"github.com/pranotoism/football-go/dto"
	"github.com/pranotoism/football-go/model"
	"github.com/pranotoism/football-go/repository"
	"gorm.io/gorm"
)

type MatchService struct {
	matchRepo  *repository.MatchRepository
	teamRepo   *repository.TeamRepository
	goalRepo   *repository.GoalRepository
}

func NewMatchService(matchRepo *repository.MatchRepository, teamRepo *repository.TeamRepository, goalRepo *repository.GoalRepository) *MatchService {
	return &MatchService{matchRepo: matchRepo, teamRepo: teamRepo, goalRepo: goalRepo}
}

func (s *MatchService) Create(req dto.CreateMatchRequest) (*model.Match, error) {
	if req.HomeTeamID == req.AwayTeamID {
		return nil, errors.New("home team and away team cannot be the same")
	}

	if !s.teamRepo.Exists(req.HomeTeamID) {
		return nil, errors.New("home team not found")
	}
	if !s.teamRepo.Exists(req.AwayTeamID) {
		return nil, errors.New("away team not found")
	}

	match := &model.Match{
		MatchDate:  req.MatchDate,
		MatchTime:  req.MatchTime,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
	}

	if err := s.matchRepo.Create(match); err != nil {
		return nil, err
	}

	return s.matchRepo.FindByID(match.ID)
}

func (s *MatchService) FindAll(page, perPage int) ([]model.Match, int64, error) {
	return s.matchRepo.FindAll(page, perPage)
}

func (s *MatchService) FindByID(id uint) (*model.Match, error) {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}
	return match, nil
}

func (s *MatchService) Update(id uint, req dto.UpdateMatchRequest) (*model.Match, error) {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	if req.HomeTeamID != 0 && req.AwayTeamID != 0 && req.HomeTeamID == req.AwayTeamID {
		return nil, errors.New("home team and away team cannot be the same")
	}

	if req.MatchDate != "" {
		match.MatchDate = req.MatchDate
	}
	if req.MatchTime != "" {
		match.MatchTime = req.MatchTime
	}
	if req.HomeTeamID != 0 {
		if !s.teamRepo.Exists(req.HomeTeamID) {
			return nil, errors.New("home team not found")
		}
		match.HomeTeamID = req.HomeTeamID
	}
	if req.AwayTeamID != 0 {
		if !s.teamRepo.Exists(req.AwayTeamID) {
			return nil, errors.New("away team not found")
		}
		match.AwayTeamID = req.AwayTeamID
	}

	if err := s.matchRepo.Update(match); err != nil {
		return nil, err
	}

	return s.matchRepo.FindByID(match.ID)
}

func (s *MatchService) Delete(id uint) error {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("match not found")
		}
		return err
	}

	// Cascade soft-delete related goals
	if err := s.goalRepo.DeleteByMatchID(id); err != nil {
		return err
	}

	return s.matchRepo.Delete(match)
}

func (s *MatchService) ReportResult(id uint, req dto.ReportResultRequest) (*model.Match, error) {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	if match.HomeScore != nil {
		return nil, errors.New("match result already reported")
	}

	// Validate goal counts per team
	homeGoals := 0
	awayGoals := 0
	for _, g := range req.Goals {
		if g.TeamID == match.HomeTeamID {
			homeGoals++
		} else if g.TeamID == match.AwayTeamID {
			awayGoals++
		} else {
			return nil, fmt.Errorf("goal team_id %d does not belong to either team in this match", g.TeamID)
		}
	}

	if homeGoals != req.HomeScore {
		return nil, fmt.Errorf("home goal count (%d) does not match home_score (%d)", homeGoals, req.HomeScore)
	}
	if awayGoals != req.AwayScore {
		return nil, fmt.Errorf("away goal count (%d) does not match away_score (%d)", awayGoals, req.AwayScore)
	}

	// Transaction: update scores + create goals
	db := s.matchRepo.DB()
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Match{}).Where("id = ?", id).Updates(map[string]interface{}{
			"home_score": req.HomeScore,
			"away_score": req.AwayScore,
		}).Error; err != nil {
			return err
		}

		if len(req.Goals) > 0 {
			goals := make([]model.Goal, len(req.Goals))
			for i, g := range req.Goals {
				goals[i] = model.Goal{
					MatchID:  id,
					PlayerID: g.PlayerID,
					TeamID:   g.TeamID,
					Minute:   g.Minute,
				}
			}
			if err := tx.Create(&goals).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.matchRepo.FindByID(id)
}
