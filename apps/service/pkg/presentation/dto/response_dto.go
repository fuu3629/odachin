package dto

import (
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
)

func ToUserInfoResponse(u *models.User) *odachin.GetUserInfoResponse {
	return &odachin.GetUserInfoResponse{
		UserId:         u.UserID,
		Name:           u.UserName,
		Role:           odachin.Role(odachin.Role_value[u.Role]),
		AvatarImageUrl: u.AvatarImageUrl,
	}
}

func ToOwnInfoResponse(u *models.User) *odachin.GetOwnInfoResponse {
	return &odachin.GetOwnInfoResponse{
		Name:           u.UserName,
		Email:          u.Email,
		AvaterImageUrl: u.AvatarImageUrl,
	}
}

func ToGetRewardListResponse(r []models.RewardPeriod) *odachin.GetRewardListResponse {
	rewardList := make([]*odachin.RewardInfo, len(r))
	for i, reward := range r {
		rewardList[i] = &odachin.RewardInfo{
			RewardPeriodId: uint64(reward.RewardPeriodID),
			FromUserId:     reward.Reward.FromUserID,
			ToUserId:       reward.Reward.ToUserID,
			Amount:         reward.Reward.Amount,
			RewardType:     odachin.Reward_Type(odachin.Reward_Type_value[reward.Reward.PeriodType]),
			Title:          reward.Reward.Title,
			Description:    reward.Reward.Description,
			IsCompleted:    reward.IsCompleted,
			IsEditable:     reward.IsEditable,
		}
	}
	return &odachin.GetRewardListResponse{
		RewardList: rewardList,
	}
}
