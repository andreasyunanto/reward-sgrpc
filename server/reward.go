package server

import (
	"context"

	pbr "github.com/andreasyunanto/reward-sgrpc/pb"
	"github.com/andreasyunanto/reward-sgrpc/repositories"
	"github.com/andreasyunanto/reward-sgrpc/services"
)

type RewardServer struct{}

// Get Reward By User ID
func (s *RewardServer) GetRewardUser(ctx context.Context, req *pbr.GetRewardUserRequest) (*pbr.GetRewardResponse, error) {
	return services.GetRewardUser(&repositories.RewardRepository{}, req.UserId)
}
