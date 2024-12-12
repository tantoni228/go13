package models

type Role struct {
	Id                int    `db:"id"`
	Name              string `db:"name"`
	IsSystem          bool   `db:"is_system"`
	CanManageMembers  bool   `db:"can_manage_members"`
	CanEditRoles      bool   `db:"can_edit_roles"`
	CanDeleteMessages bool   `db:"can_delete_messages"`
	CanGetJoinCode    bool   `db:"can_get_join_code"`
	CanEditChatInfo   bool   `db:"can_edit_chat_info"`
	CanDeleteChat     bool   `db:"can_delete_chat"`
}

var (
	RoleMember = Role{
		Name:     "member",
		IsSystem: true,
	}

	RoleCreator = Role{
		Name:              "creator",
		IsSystem:          true,
		CanManageMembers:  true,
		CanEditRoles:      true,
		CanDeleteMessages: true,
		CanGetJoinCode:    true,
		CanEditChatInfo:   true,
		CanDeleteChat:     true,
	}
)
