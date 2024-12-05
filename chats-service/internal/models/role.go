package models

type Role struct {
	Id                int
	Name              string
	CanBanUsers       bool
	CanEditRoles      bool
	CanDeleteMessages bool
	CanGetJoinCode    bool
	CanEditChatInfo   bool
	CanDeleteChat     bool
}
