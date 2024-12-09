// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

// BanUserNoContent is response for BanUser operation.
type BanUserNoContent struct{}

func (*BanUserNoContent) banUserRes() {}

// BanUserNotFound is response for BanUser operation.
type BanUserNotFound struct{}

func (*BanUserNotFound) banUserRes() {}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/chat
type Chat struct {
	ID          ChatId `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetID returns the value of ID.
func (s *Chat) GetID() ChatId {
	return s.ID
}

// GetName returns the value of Name.
func (s *Chat) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Chat) GetDescription() string {
	return s.Description
}

// SetID sets the value of ID.
func (s *Chat) SetID(val ChatId) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Chat) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Chat) SetDescription(val string) {
	s.Description = val
}

func (*Chat) createChatRes()  {}
func (*Chat) getChatByIdRes() {}
func (*Chat) updateChatRes()  {}

type ChatId int

// Ref: #/components/schemas/chatInput
type ChatInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetName returns the value of Name.
func (s *ChatInput) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *ChatInput) GetDescription() string {
	return s.Description
}

// SetName sets the value of Name.
func (s *ChatInput) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *ChatInput) SetDescription(val string) {
	s.Description = val
}

// Ref: #/components/responses/chatNotFoundResponse
type ChatNotFoundResponse struct{}

func (*ChatNotFoundResponse) createRoleRes()  {}
func (*ChatNotFoundResponse) deleteChatRes()  {}
func (*ChatNotFoundResponse) getChatByIdRes() {}
func (*ChatNotFoundResponse) getJoinCodeRes() {}
func (*ChatNotFoundResponse) joinChatRes()    {}
func (*ChatNotFoundResponse) leaveChatRes()   {}
func (*ChatNotFoundResponse) listMembersRes() {}
func (*ChatNotFoundResponse) listRolesRes()   {}
func (*ChatNotFoundResponse) updateChatRes()  {}

// CheckAccessNoContent is response for CheckAccess operation.
type CheckAccessNoContent struct{}

func (*CheckAccessNoContent) checkAccessRes() {}

// CheckAccessNotFound is response for CheckAccess operation.
type CheckAccessNotFound struct{}

func (*CheckAccessNotFound) checkAccessRes() {}

type CheckAccessXTargetMethod string

const (
	CheckAccessXTargetMethodGET    CheckAccessXTargetMethod = "GET"
	CheckAccessXTargetMethodPOST   CheckAccessXTargetMethod = "POST"
	CheckAccessXTargetMethodPUT    CheckAccessXTargetMethod = "PUT"
	CheckAccessXTargetMethodDELETE CheckAccessXTargetMethod = "DELETE"
)

// AllValues returns all CheckAccessXTargetMethod values.
func (CheckAccessXTargetMethod) AllValues() []CheckAccessXTargetMethod {
	return []CheckAccessXTargetMethod{
		CheckAccessXTargetMethodGET,
		CheckAccessXTargetMethodPOST,
		CheckAccessXTargetMethodPUT,
		CheckAccessXTargetMethodDELETE,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CheckAccessXTargetMethod) MarshalText() ([]byte, error) {
	switch s {
	case CheckAccessXTargetMethodGET:
		return []byte(s), nil
	case CheckAccessXTargetMethodPOST:
		return []byte(s), nil
	case CheckAccessXTargetMethodPUT:
		return []byte(s), nil
	case CheckAccessXTargetMethodDELETE:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CheckAccessXTargetMethod) UnmarshalText(data []byte) error {
	switch CheckAccessXTargetMethod(data) {
	case CheckAccessXTargetMethodGET:
		*s = CheckAccessXTargetMethodGET
		return nil
	case CheckAccessXTargetMethodPOST:
		*s = CheckAccessXTargetMethodPOST
		return nil
	case CheckAccessXTargetMethodPUT:
		*s = CheckAccessXTargetMethodPUT
		return nil
	case CheckAccessXTargetMethodDELETE:
		*s = CheckAccessXTargetMethodDELETE
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// CreateRoleConflict is response for CreateRole operation.
type CreateRoleConflict struct{}

func (*CreateRoleConflict) createRoleRes() {}

// DeleteChatNoContent is response for DeleteChat operation.
type DeleteChatNoContent struct{}

func (*DeleteChatNoContent) deleteChatRes() {}

// DeleteRoleNoContent is response for DeleteRole operation.
type DeleteRoleNoContent struct{}

func (*DeleteRoleNoContent) deleteRoleRes() {}

// DeleteRoleNotFound is response for DeleteRole operation.
type DeleteRoleNotFound struct{}

func (*DeleteRoleNotFound) deleteRoleRes() {}

// GetRoleByIdNotFound is response for GetRoleById operation.
type GetRoleByIdNotFound struct{}

func (*GetRoleByIdNotFound) getRoleByIdRes() {}

// Ref: #/components/responses/internalErrorResponse
type InternalErrorResponse struct{}

func (*InternalErrorResponse) banUserRes()     {}
func (*InternalErrorResponse) checkAccessRes() {}
func (*InternalErrorResponse) createChatRes()  {}
func (*InternalErrorResponse) createRoleRes()  {}
func (*InternalErrorResponse) deleteChatRes()  {}
func (*InternalErrorResponse) deleteRoleRes()  {}
func (*InternalErrorResponse) getChatByIdRes() {}
func (*InternalErrorResponse) getJoinCodeRes() {}
func (*InternalErrorResponse) getRoleByIdRes() {}
func (*InternalErrorResponse) joinChatRes()    {}
func (*InternalErrorResponse) leaveChatRes()   {}
func (*InternalErrorResponse) listChatsRes()   {}
func (*InternalErrorResponse) listMembersRes() {}
func (*InternalErrorResponse) listRolesRes()   {}
func (*InternalErrorResponse) setRoleRes()     {}
func (*InternalErrorResponse) updateChatRes()  {}
func (*InternalErrorResponse) updateRoleRes()  {}

type InvalidInputResponse struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *InvalidInputResponse) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *InvalidInputResponse) SetMessage(val string) {
	s.Message = val
}

func (*InvalidInputResponse) banUserRes()     {}
func (*InvalidInputResponse) checkAccessRes() {}
func (*InvalidInputResponse) createChatRes()  {}
func (*InvalidInputResponse) createRoleRes()  {}
func (*InvalidInputResponse) deleteChatRes()  {}
func (*InvalidInputResponse) deleteRoleRes()  {}
func (*InvalidInputResponse) getChatByIdRes() {}
func (*InvalidInputResponse) getJoinCodeRes() {}
func (*InvalidInputResponse) getRoleByIdRes() {}
func (*InvalidInputResponse) joinChatRes()    {}
func (*InvalidInputResponse) leaveChatRes()   {}
func (*InvalidInputResponse) listMembersRes() {}
func (*InvalidInputResponse) listRolesRes()   {}
func (*InvalidInputResponse) setRoleRes()     {}
func (*InvalidInputResponse) updateChatRes()  {}
func (*InvalidInputResponse) updateRoleRes()  {}

// JoinChatConflict is response for JoinChat operation.
type JoinChatConflict struct{}

func (*JoinChatConflict) joinChatRes() {}

// JoinChatNoContent is response for JoinChat operation.
type JoinChatNoContent struct{}

func (*JoinChatNoContent) joinChatRes() {}

type JoinChatReq struct {
	JoinCode string `json:"join_code"`
}

// GetJoinCode returns the value of JoinCode.
func (s *JoinChatReq) GetJoinCode() string {
	return s.JoinCode
}

// SetJoinCode sets the value of JoinCode.
func (s *JoinChatReq) SetJoinCode(val string) {
	s.JoinCode = val
}

type JoinCodeResponse struct {
	JoinCode string `json:"join_code"`
}

// GetJoinCode returns the value of JoinCode.
func (s *JoinCodeResponse) GetJoinCode() string {
	return s.JoinCode
}

// SetJoinCode sets the value of JoinCode.
func (s *JoinCodeResponse) SetJoinCode(val string) {
	s.JoinCode = val
}

func (*JoinCodeResponse) getJoinCodeRes() {}

// LeaveChatNoContent is response for LeaveChat operation.
type LeaveChatNoContent struct{}

func (*LeaveChatNoContent) leaveChatRes() {}

type ListChatsOKApplicationJSON []Chat

func (*ListChatsOKApplicationJSON) listChatsRes() {}

type ListMembersOKApplicationJSON []Member

func (*ListMembersOKApplicationJSON) listMembersRes() {}

type ListRolesOKApplicationJSON []Role

func (*ListRolesOKApplicationJSON) listRolesRes() {}

// Ref: #/components/schemas/member
type Member struct {
	UserID UserId `json:"user_id"`
	RoleID RoleId `json:"role_id"`
}

// GetUserID returns the value of UserID.
func (s *Member) GetUserID() UserId {
	return s.UserID
}

// GetRoleID returns the value of RoleID.
func (s *Member) GetRoleID() RoleId {
	return s.RoleID
}

// SetUserID sets the value of UserID.
func (s *Member) SetUserID(val UserId) {
	s.UserID = val
}

// SetRoleID sets the value of RoleID.
func (s *Member) SetRoleID(val RoleId) {
	s.RoleID = val
}

// Ref: #/components/schemas/role
type Role struct {
	ID                RoleId `json:"id"`
	Name              string `json:"name"`
	IsSystem          bool   `json:"is_system"`
	CanBanUsers       bool   `json:"can_ban_users"`
	CanEditRoles      bool   `json:"can_edit_roles"`
	CanDeleteMessages bool   `json:"can_delete_messages"`
	CanGetJoinCode    bool   `json:"can_get_join_code"`
	CanEditChatInfo   bool   `json:"can_edit_chat_info"`
	CanDeleteChat     bool   `json:"can_delete_chat"`
}

// GetID returns the value of ID.
func (s *Role) GetID() RoleId {
	return s.ID
}

// GetName returns the value of Name.
func (s *Role) GetName() string {
	return s.Name
}

// GetIsSystem returns the value of IsSystem.
func (s *Role) GetIsSystem() bool {
	return s.IsSystem
}

// GetCanBanUsers returns the value of CanBanUsers.
func (s *Role) GetCanBanUsers() bool {
	return s.CanBanUsers
}

// GetCanEditRoles returns the value of CanEditRoles.
func (s *Role) GetCanEditRoles() bool {
	return s.CanEditRoles
}

// GetCanDeleteMessages returns the value of CanDeleteMessages.
func (s *Role) GetCanDeleteMessages() bool {
	return s.CanDeleteMessages
}

// GetCanGetJoinCode returns the value of CanGetJoinCode.
func (s *Role) GetCanGetJoinCode() bool {
	return s.CanGetJoinCode
}

// GetCanEditChatInfo returns the value of CanEditChatInfo.
func (s *Role) GetCanEditChatInfo() bool {
	return s.CanEditChatInfo
}

// GetCanDeleteChat returns the value of CanDeleteChat.
func (s *Role) GetCanDeleteChat() bool {
	return s.CanDeleteChat
}

// SetID sets the value of ID.
func (s *Role) SetID(val RoleId) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Role) SetName(val string) {
	s.Name = val
}

// SetIsSystem sets the value of IsSystem.
func (s *Role) SetIsSystem(val bool) {
	s.IsSystem = val
}

// SetCanBanUsers sets the value of CanBanUsers.
func (s *Role) SetCanBanUsers(val bool) {
	s.CanBanUsers = val
}

// SetCanEditRoles sets the value of CanEditRoles.
func (s *Role) SetCanEditRoles(val bool) {
	s.CanEditRoles = val
}

// SetCanDeleteMessages sets the value of CanDeleteMessages.
func (s *Role) SetCanDeleteMessages(val bool) {
	s.CanDeleteMessages = val
}

// SetCanGetJoinCode sets the value of CanGetJoinCode.
func (s *Role) SetCanGetJoinCode(val bool) {
	s.CanGetJoinCode = val
}

// SetCanEditChatInfo sets the value of CanEditChatInfo.
func (s *Role) SetCanEditChatInfo(val bool) {
	s.CanEditChatInfo = val
}

// SetCanDeleteChat sets the value of CanDeleteChat.
func (s *Role) SetCanDeleteChat(val bool) {
	s.CanDeleteChat = val
}

func (*Role) createRoleRes()  {}
func (*Role) getRoleByIdRes() {}
func (*Role) updateRoleRes()  {}

type RoleId int

// Ref: #/components/schemas/roleInput
type RoleInput struct {
	Name              string `json:"name"`
	CanBanUsers       bool   `json:"can_ban_users"`
	CanEditRoles      bool   `json:"can_edit_roles"`
	CanDeleteMessages bool   `json:"can_delete_messages"`
	CanGetJoinCode    bool   `json:"can_get_join_code"`
	CanEditChatInfo   bool   `json:"can_edit_chat_info"`
	CanDeleteChat     bool   `json:"can_delete_chat"`
}

// GetName returns the value of Name.
func (s *RoleInput) GetName() string {
	return s.Name
}

// GetCanBanUsers returns the value of CanBanUsers.
func (s *RoleInput) GetCanBanUsers() bool {
	return s.CanBanUsers
}

// GetCanEditRoles returns the value of CanEditRoles.
func (s *RoleInput) GetCanEditRoles() bool {
	return s.CanEditRoles
}

// GetCanDeleteMessages returns the value of CanDeleteMessages.
func (s *RoleInput) GetCanDeleteMessages() bool {
	return s.CanDeleteMessages
}

// GetCanGetJoinCode returns the value of CanGetJoinCode.
func (s *RoleInput) GetCanGetJoinCode() bool {
	return s.CanGetJoinCode
}

// GetCanEditChatInfo returns the value of CanEditChatInfo.
func (s *RoleInput) GetCanEditChatInfo() bool {
	return s.CanEditChatInfo
}

// GetCanDeleteChat returns the value of CanDeleteChat.
func (s *RoleInput) GetCanDeleteChat() bool {
	return s.CanDeleteChat
}

// SetName sets the value of Name.
func (s *RoleInput) SetName(val string) {
	s.Name = val
}

// SetCanBanUsers sets the value of CanBanUsers.
func (s *RoleInput) SetCanBanUsers(val bool) {
	s.CanBanUsers = val
}

// SetCanEditRoles sets the value of CanEditRoles.
func (s *RoleInput) SetCanEditRoles(val bool) {
	s.CanEditRoles = val
}

// SetCanDeleteMessages sets the value of CanDeleteMessages.
func (s *RoleInput) SetCanDeleteMessages(val bool) {
	s.CanDeleteMessages = val
}

// SetCanGetJoinCode sets the value of CanGetJoinCode.
func (s *RoleInput) SetCanGetJoinCode(val bool) {
	s.CanGetJoinCode = val
}

// SetCanEditChatInfo sets the value of CanEditChatInfo.
func (s *RoleInput) SetCanEditChatInfo(val bool) {
	s.CanEditChatInfo = val
}

// SetCanDeleteChat sets the value of CanDeleteChat.
func (s *RoleInput) SetCanDeleteChat(val bool) {
	s.CanDeleteChat = val
}

// SetRoleNoContent is response for SetRole operation.
type SetRoleNoContent struct{}

func (*SetRoleNoContent) setRoleRes() {}

// SetRoleNotFound is response for SetRole operation.
type SetRoleNotFound struct{}

func (*SetRoleNotFound) setRoleRes() {}

type SetRoleReq struct {
	RoleID RoleId `json:"role_id"`
}

// GetRoleID returns the value of RoleID.
func (s *SetRoleReq) GetRoleID() RoleId {
	return s.RoleID
}

// SetRoleID sets the value of RoleID.
func (s *SetRoleReq) SetRoleID(val RoleId) {
	s.RoleID = val
}

// Ref: #/components/responses/unauthenticatedResponse
type UnauthenticatedResponse struct{}

func (*UnauthenticatedResponse) banUserRes()     {}
func (*UnauthenticatedResponse) checkAccessRes() {}
func (*UnauthenticatedResponse) createChatRes()  {}
func (*UnauthenticatedResponse) createRoleRes()  {}
func (*UnauthenticatedResponse) deleteChatRes()  {}
func (*UnauthenticatedResponse) deleteRoleRes()  {}
func (*UnauthenticatedResponse) getChatByIdRes() {}
func (*UnauthenticatedResponse) getJoinCodeRes() {}
func (*UnauthenticatedResponse) getRoleByIdRes() {}
func (*UnauthenticatedResponse) joinChatRes()    {}
func (*UnauthenticatedResponse) leaveChatRes()   {}
func (*UnauthenticatedResponse) listChatsRes()   {}
func (*UnauthenticatedResponse) listMembersRes() {}
func (*UnauthenticatedResponse) listRolesRes()   {}
func (*UnauthenticatedResponse) setRoleRes()     {}
func (*UnauthenticatedResponse) updateChatRes()  {}
func (*UnauthenticatedResponse) updateRoleRes()  {}

// Ref: #/components/responses/unauthorizedResponse
type UnauthorizedResponse struct{}

func (*UnauthorizedResponse) banUserRes()     {}
func (*UnauthorizedResponse) checkAccessRes() {}
func (*UnauthorizedResponse) createRoleRes()  {}
func (*UnauthorizedResponse) deleteChatRes()  {}
func (*UnauthorizedResponse) deleteRoleRes()  {}
func (*UnauthorizedResponse) getChatByIdRes() {}
func (*UnauthorizedResponse) getJoinCodeRes() {}
func (*UnauthorizedResponse) getRoleByIdRes() {}
func (*UnauthorizedResponse) joinChatRes()    {}
func (*UnauthorizedResponse) listMembersRes() {}
func (*UnauthorizedResponse) listRolesRes()   {}
func (*UnauthorizedResponse) setRoleRes()     {}
func (*UnauthorizedResponse) updateChatRes()  {}
func (*UnauthorizedResponse) updateRoleRes()  {}

// UpdateRoleConflict is response for UpdateRole operation.
type UpdateRoleConflict struct{}

func (*UpdateRoleConflict) updateRoleRes() {}

// UpdateRoleNotFound is response for UpdateRole operation.
type UpdateRoleNotFound struct{}

func (*UpdateRoleNotFound) updateRoleRes() {}

type UserId string
