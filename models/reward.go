package models

import "time"

type Reward struct {
	RewardId   string    `gorm:"primaryKey;type:varchar(36);" json:"rewardId"`
	UserId     string    `gorm:"type:varchar(36);index;" json:"userId"`
	TotalPoint int       `gorm:"default:0;type:bigint;" json:"totalPoint"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Rewards []Reward
