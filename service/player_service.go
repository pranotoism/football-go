package service

import (
	"errors"

	"github.com/pranotoism/football-go/dto"
	"github.com/pranotoism/football-go/model"
	"github.com/pranotoism/football-go/repository"
	"gorm.io/gorm"
)

type PlayerService struct {
	playerRepo *repository.PlayerRepository
	teamRepo   *repository.TeamRepository
}

func NewPlayerService(playerRepo *repository.PlayerRepository, teamRepo *repository.TeamRepository) *PlayerService {
	return &PlayerService{playerRepo: playerRepo, teamRepo: teamRepo}
}

func (s *PlayerService) Create(teamID uint, req dto.CreatePlayerRequest) (*model.Player, error) {
	if !s.teamRepo.Exists(teamID) {
		return nil, errors.New("team not found")
	}

	if s.playerRepo.IsJerseyNumberTaken(teamID, req.JerseyNumber, 0) {
		return nil, errors.New("jersey number already taken in this team")
	}

	player := &model.Player{
		TeamID:       teamID,
		Name:         req.Name,
		HeightCM:     req.HeightCM,
		WeightKG:     req.WeightKG,
		Position:     req.Position,
		JerseyNumber: req.JerseyNumber,
	}

	if err := s.playerRepo.Create(player); err != nil {
		return nil, err
	}
	return s.playerRepo.FindByID(player.ID)
}

func (s *PlayerService) FindByTeam(teamID uint, page, perPage int) ([]model.Player, int64, error) {
	if !s.teamRepo.Exists(teamID) {
		return nil, 0, errors.New("team not found")
	}
	return s.playerRepo.FindByTeam(teamID, page, perPage)
}

func (s *PlayerService) FindByID(id uint) (*model.Player, error) {
	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player not found")
		}
		return nil, err
	}
	return player, nil
}

func (s *PlayerService) Update(id uint, req dto.UpdatePlayerRequest) (*model.Player, error) {
	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player not found")
		}
		return nil, err
	}

	if req.JerseyNumber != 0 && req.JerseyNumber != player.JerseyNumber {
		if s.playerRepo.IsJerseyNumberTaken(player.TeamID, req.JerseyNumber, id) {
			return nil, errors.New("jersey number already taken in this team")
		}
		player.JerseyNumber = req.JerseyNumber
	}

	if req.Name != "" {
		player.Name = req.Name
	}
	if req.HeightCM != 0 {
		player.HeightCM = req.HeightCM
	}
	if req.WeightKG != 0 {
		player.WeightKG = req.WeightKG
	}
	if req.Position != "" {
		player.Position = req.Position
	}

	if err := s.playerRepo.Update(player); err != nil {
		return nil, err
	}
	return player, nil
}

func (s *PlayerService) Delete(id uint) error {
	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("player not found")
		}
		return err
	}
	return s.playerRepo.Delete(player)
}
