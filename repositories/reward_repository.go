package repositories

import (
	"github.com/andreasyunanto/reward-sgrpc/database"
	"github.com/andreasyunanto/reward-sgrpc/models"
	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{db: db}
}

// Get Post By Id
func (r *RewardRepository) GetRewardByUser(id string) (models.Reward, error) {

	var reward models.Reward

	err := database.DB.Where("user_id = ?", id).First(&reward).Error
	if err != nil {
		return reward, err
	}

	return reward, nil

}
