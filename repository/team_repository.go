package repository

import (
	"github.com/pranotoism/football-go/model"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Create(team *model.Team) error {
	return r.db.Create(team).Error
}

func (r *TeamRepository) FindAll(page, perPage int) ([]model.Team, int64, error) {
	var teams []model.Team
	var total int64

	r.db.Model(&model.Team{}).Count(&total)

	offset := (page - 1) * perPage
	err := r.db.Offset(offset).Limit(perPage).Find(&teams).Error
	return teams, total, err
}

func (r *TeamRepository) FindByID(id uint) (*model.Team, error) {
	var team model.Team
	err := r.db.Preload("Players").First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepository) Update(team *model.Team) error {
	return r.db.Save(team).Error
}

func (r *TeamRepository) Delete(team *model.Team) error {
	return r.db.Delete(team).Error
}

func (r *TeamRepository) Exists(id uint) bool {
	var count int64
	r.db.Model(&model.Team{}).Where("id = ?", id).Count(&count)
	return count > 0
}
