package repository

import (
	"github.com/pranotoism/football-go/model"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(player *model.Player) error {
	return r.db.Create(player).Error
}

func (r *PlayerRepository) FindByTeam(teamID uint, page, perPage int) ([]model.Player, int64, error) {
	var players []model.Player
	var total int64

	r.db.Model(&model.Player{}).Where("team_id = ?", teamID).Count(&total)

	offset := (page - 1) * perPage
	err := r.db.Where("team_id = ?", teamID).Offset(offset).Limit(perPage).Find(&players).Error
	return players, total, err
}

func (r *PlayerRepository) FindByID(id uint) (*model.Player, error) {
	var player model.Player
	err := r.db.Preload("Team").First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) Update(player *model.Player) error {
	return r.db.Save(player).Error
}

func (r *PlayerRepository) Delete(player *model.Player) error {
	return r.db.Delete(player).Error
}

func (r *PlayerRepository) IsJerseyNumberTaken(teamID uint, jerseyNumber int, excludePlayerID uint) bool {
	var count int64
	query := r.db.Model(&model.Player{}).Where("team_id = ? AND jersey_number = ?", teamID, jerseyNumber)
	if excludePlayerID > 0 {
		query = query.Where("id != ?", excludePlayerID)
	}
	query.Count(&count)
	return count > 0
}

func (r *PlayerRepository) SoftDeleteByTeamID(teamID uint) error {
	return r.db.Where("team_id = ?", teamID).Delete(&model.Player{}).Error
}
