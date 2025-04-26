package dto

import (
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/internal/models"
)

func uintToUint64Pointer(value *uint) *uint64 {
	if value == nil {
		return nil
	}
	converted := uint64(*value)
	return &converted
}

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
		Role:           odachin.Role(odachin.Role_value[u.Role]),
		FamilyId:       uintToUint64Pointer(u.FamilyID),
		UserId:         u.UserID,
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

func ToGetFamilyInfoResponse(members []models.User, family *models.Family) *odachin.GetFamilyInfoResponse {
	familyMembers := make([]*odachin.FamilyUser, len(members))
	for i, member := range members {
		familyMembers[i] = &odachin.FamilyUser{
			UserId:         member.UserID,
			Name:           member.UserName,
			Role:           odachin.Role(odachin.Role_value[member.Role]),
			AvatarImageUrl: member.AvatarImageUrl,
		}
	}
	return &odachin.GetFamilyInfoResponse{
		FamilyId:      uint64(family.FamilyID),
		FamilyName:    family.FamilyName,
		FamilyMembers: familyMembers,
	}
}

func ToGetAllowanceByFromUserIdResponse(allowanceList []models.Allowance, userList []models.User) *odachin.GetAllowanceByFromUserIdResponse {
	allowances := make([]*odachin.Allowance, len(allowanceList))
	for i, allowance := range allowanceList {
		var dayOfWeek *odachin.DayOfWeek
		if allowance.DayOfWeek != nil {
			dayOfWeekValue := odachin.DayOfWeek(odachin.DayOfWeek_value[*allowance.DayOfWeek])
			dayOfWeek = &dayOfWeekValue
		}
		allowances[i] = &odachin.Allowance{
			AllowanceId:    uint64(allowance.AllowanceID),
			ToUserId:       allowance.ToUserID,
			ToUserName:     userList[i].UserName,
			Amount:         allowance.Amount,
			IntervalType:   odachin.Alloance_Type(odachin.Alloance_Type_value[allowance.IntervalType]),
			Date:           allowance.Date,
			DayOfWeek:      dayOfWeek,
			AvatarImageUrl: userList[i].AvatarImageUrl,
		}
	}
	return &odachin.GetAllowanceByFromUserIdResponse{
		Allowances: allowances,
	}
}
