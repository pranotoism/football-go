package service

import (
	"errors"

	"github.com/pranotoism/football-go/dto"
	"github.com/pranotoism/football-go/model"
	"github.com/pranotoism/football-go/repository"
	"gorm.io/gorm"
)

type TeamService struct {
	teamRepo   *repository.TeamRepository
	playerRepo *repository.PlayerRepository
}

func NewTeamService(teamRepo *repository.TeamRepository, playerRepo *repository.PlayerRepository) *TeamService {
	return &TeamService{teamRepo: teamRepo, playerRepo: playerRepo}
}

func (s *TeamService) Create(req dto.CreateTeamRequest) (*model.Team, error) {
	team := &model.Team{
		Name:        req.Name,
		LogoURL:     req.LogoURL,
		FoundedYear: req.FoundedYear,
		HQAddress:   req.HQAddress,
		HQCity:      req.HQCity,
	}

	if err := s.teamRepo.Create(team); err != nil {
		return nil, err
	}
	return team, nil
}

func (s *TeamService) FindAll(page, perPage int) ([]model.Team, int64, error) {
	return s.teamRepo.FindAll(page, perPage)
}

func (s *TeamService) FindByID(id uint) (*model.Team, error) {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}
	return team, nil
}

func (s *TeamService) Update(id uint, req dto.UpdateTeamRequest) (*model.Team, error) {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}

	if req.Name != "" {
		team.Name = req.Name
	}
	if req.LogoURL != "" {
		team.LogoURL = req.LogoURL
	}
	if req.FoundedYear != 0 {
		team.FoundedYear = req.FoundedYear
	}
	if req.HQAddress != "" {
		team.HQAddress = req.HQAddress
	}
	if req.HQCity != "" {
		team.HQCity = req.HQCity
	}

	if err := s.teamRepo.Update(team); err != nil {
		return nil, err
	}
	return team, nil
}

func (s *TeamService) Delete(id uint) error {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("team not found")
		}
		return err
	}

	// Cascade soft-delete players
	if err := s.playerRepo.SoftDeleteByTeamID(id); err != nil {
		return err
	}

	return s.teamRepo.Delete(team)
}
