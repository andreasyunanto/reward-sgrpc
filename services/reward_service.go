package services

import (
	"errors"

	"github.com/andreasyunanto/reward-sgrpc/pb"
	pbr "github.com/andreasyunanto/reward-sgrpc/pb"
	"github.com/andreasyunanto/reward-sgrpc/repositories"
)

func GetRewardUser(repo *repositories.RewardRepository, id string) (*pbr.GetRewardResponse, error) {
	operationResult, err := repo.GetRewardByUser(id)

	if err != nil {
		return &pb.GetRewardResponse{Status: false, Data: nil, Message: err.Error()}, errors.New(err.Error())
	}

	var data *pbr.Reward = &pbr.Reward{
		UserId:     operationResult.UserId,
		RewardId:   operationResult.RewardId,
		TotalPoint: int64(operationResult.TotalPoint),
	}

	return &pbr.GetRewardResponse{
		Message: "OK",
		Data:    data,
		Status:  true,
	}, nil
}
