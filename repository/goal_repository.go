package repository

import (
	"github.com/pranotoism/football-go/model"
	"gorm.io/gorm"
)

type GoalRepository struct {
	db *gorm.DB
}

func NewGoalRepository(db *gorm.DB) *GoalRepository {
	return &GoalRepository{db: db}
}

func (r *GoalRepository) CreateBatch(goals []model.Goal) error {
	if len(goals) == 0 {
		return nil
	}
	return r.db.Create(&goals).Error
}

func (r *GoalRepository) CreateBatchTx(tx *gorm.DB, goals []model.Goal) error {
	if len(goals) == 0 {
		return nil
	}
	return tx.Create(&goals).Error
}

func (r *GoalRepository) FindByMatchID(matchID uint) ([]model.Goal, error) {
	var goals []model.Goal
	err := r.db.Where("match_id = ?", matchID).
		Preload("Player").Preload("Team").
		Order("minute ASC").Find(&goals).Error
	return goals, err
}

func (r *GoalRepository) DeleteByMatchID(matchID uint) error {
	return r.db.Where("match_id = ?", matchID).Delete(&model.Goal{}).Error
}
