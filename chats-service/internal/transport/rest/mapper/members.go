package mapper

import (
	"go13/chats-service/internal/models"
	api "go13/pkg/ogen/chats-service"
)

func ModelsMemberToApiMember(member models.Member) *api.Member {
	return &api.Member{
		UserID: api.UserId(member.UserId),
		RoleID: api.RoleId(member.RoleId),
	}
}
