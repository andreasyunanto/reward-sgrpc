package database

import (
	"github.com/andreasyunanto/reward-sgrpc/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&models.Reward{},
	)
}
