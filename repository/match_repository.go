package repository

import (
	"github.com/pranotoism/football-go/model"
	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Create(match *model.Match) error {
	return r.db.Create(match).Error
}

func (r *MatchRepository) FindAll(page, perPage int) ([]model.Match, int64, error) {
	var matches []model.Match
	var total int64

	r.db.Model(&model.Match{}).Count(&total)

	offset := (page - 1) * perPage
	err := r.db.Preload("HomeTeam").Preload("AwayTeam").
		Offset(offset).Limit(perPage).Order("match_date DESC, match_time DESC").
		Find(&matches).Error
	return matches, total, err
}

func (r *MatchRepository) FindByID(id uint) (*model.Match, error) {
	var match model.Match
	err := r.db.Preload("HomeTeam").Preload("AwayTeam").
		Preload("Goals").Preload("Goals.Player").Preload("Goals.Team").
		First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

func (r *MatchRepository) Update(match *model.Match) error {
	return r.db.Save(match).Error
}

func (r *MatchRepository) Delete(match *model.Match) error {
	return r.db.Delete(match).Error
}

func (r *MatchRepository) FindPlayedMatches(page, perPage int) ([]model.Match, int64, error) {
	var matches []model.Match
	var total int64

	r.db.Model(&model.Match{}).Where("home_score IS NOT NULL").Count(&total)

	offset := (page - 1) * perPage
	err := r.db.Where("home_score IS NOT NULL").
		Preload("HomeTeam").Preload("AwayTeam").
		Preload("Goals").Preload("Goals.Player").Preload("Goals.Team").
		Offset(offset).Limit(perPage).Order("match_date DESC, match_time DESC").
		Find(&matches).Error
	return matches, total, err
}

func (r *MatchRepository) CountWins(teamID uint) int64 {
	var count int64
	r.db.Model(&model.Match{}).
		Where("(home_team_id = ? AND home_score > away_score) OR (away_team_id = ? AND away_score > home_score)", teamID, teamID).
		Where("home_score IS NOT NULL").
		Count(&count)
	return count
}

func (r *MatchRepository) DB() *gorm.DB {
	return r.db
}
