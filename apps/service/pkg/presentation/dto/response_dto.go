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
