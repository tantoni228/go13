// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// BanUserParams is parameters of banUser operation.
type BanUserParams struct {
	// Chat id.
	ChatId ChatId
	// User id.
	UserId UserId
}

func unpackBanUserParams(packed middleware.Parameters) (params BanUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(UserId)
	}
	return params
}

func decodeBanUserParams(args [2]string, argsEscaped bool, r *http.Request) (params BanUserParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotUserIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserId = UserId(paramsDotUserIdVal)
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := params.UserId.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CheckAccessParams is parameters of CheckAccess operation.
type CheckAccessParams struct {
	// Target uri.
	XTargetURI string
	// Target method.
	XTargetMethod CheckAccessXTargetMethod
}

func unpackCheckAccessParams(packed middleware.Parameters) (params CheckAccessParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Target-Uri",
			In:   "header",
		}
		params.XTargetURI = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Target-Method",
			In:   "header",
		}
		params.XTargetMethod = packed[key].(CheckAccessXTargetMethod)
	}
	return params
}

func decodeCheckAccessParams(args [0]string, argsEscaped bool, r *http.Request) (params CheckAccessParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Target-Uri.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Target-Uri",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XTargetURI = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Target-Uri",
			In:   "header",
			Err:  err,
		}
	}
	// Decode header: X-Target-Method.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Target-Method",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XTargetMethod = CheckAccessXTargetMethod(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.XTargetMethod.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Target-Method",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// CreateRoleParams is parameters of createRole operation.
type CreateRoleParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackCreateRoleParams(packed middleware.Parameters) (params CreateRoleParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "query",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeCreateRoleParams(args [0]string, argsEscaped bool, r *http.Request) (params CreateRoleParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: chatId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "chatId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteChatParams is parameters of deleteChat operation.
type DeleteChatParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackDeleteChatParams(packed middleware.Parameters) (params DeleteChatParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeDeleteChatParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteChatParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteRoleParams is parameters of deleteRole operation.
type DeleteRoleParams struct {
	// Role id.
	RoleId RoleId
	// Chat id.
	ChatId ChatId
}

func unpackDeleteRoleParams(packed middleware.Parameters) (params DeleteRoleParams) {
	{
		key := middleware.ParameterKey{
			Name: "roleId",
			In:   "path",
		}
		params.RoleId = packed[key].(RoleId)
	}
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "query",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeDeleteRoleParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteRoleParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: roleId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "roleId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotRoleIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotRoleIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.RoleId = RoleId(paramsDotRoleIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "roleId",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: chatId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "chatId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetChatByIdParams is parameters of getChatById operation.
type GetChatByIdParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackGetChatByIdParams(packed middleware.Parameters) (params GetChatByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeGetChatByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetChatByIdParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetJoinCodeParams is parameters of getJoinCode operation.
type GetJoinCodeParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackGetJoinCodeParams(packed middleware.Parameters) (params GetJoinCodeParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeGetJoinCodeParams(args [1]string, argsEscaped bool, r *http.Request) (params GetJoinCodeParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetRoleByIdParams is parameters of getRoleById operation.
type GetRoleByIdParams struct {
	// Role id.
	RoleId RoleId
	// Chat id.
	ChatId ChatId
}

func unpackGetRoleByIdParams(packed middleware.Parameters) (params GetRoleByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "roleId",
			In:   "path",
		}
		params.RoleId = packed[key].(RoleId)
	}
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "query",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeGetRoleByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetRoleByIdParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: roleId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "roleId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotRoleIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotRoleIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.RoleId = RoleId(paramsDotRoleIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "roleId",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: chatId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "chatId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// LeaveChatParams is parameters of leaveChat operation.
type LeaveChatParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackLeaveChatParams(packed middleware.Parameters) (params LeaveChatParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeLeaveChatParams(args [1]string, argsEscaped bool, r *http.Request) (params LeaveChatParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListMembersParams is parameters of listMembers operation.
type ListMembersParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackListMembersParams(packed middleware.Parameters) (params ListMembersParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeListMembersParams(args [1]string, argsEscaped bool, r *http.Request) (params ListMembersParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListRolesParams is parameters of listRoles operation.
type ListRolesParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackListRolesParams(packed middleware.Parameters) (params ListRolesParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "query",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeListRolesParams(args [0]string, argsEscaped bool, r *http.Request) (params ListRolesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: chatId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "chatId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// SetRoleParams is parameters of setRole operation.
type SetRoleParams struct {
	// Chat id.
	ChatId ChatId
	// User id.
	UserId UserId
}

func unpackSetRoleParams(packed middleware.Parameters) (params SetRoleParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(UserId)
	}
	return params
}

func decodeSetRoleParams(args [2]string, argsEscaped bool, r *http.Request) (params SetRoleParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: userId.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotUserIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserId = UserId(paramsDotUserIdVal)
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := params.UserId.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateChatParams is parameters of updateChat operation.
type UpdateChatParams struct {
	// Chat id.
	ChatId ChatId
}

func unpackUpdateChatParams(packed middleware.Parameters) (params UpdateChatParams) {
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "path",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeUpdateChatParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateChatParams, _ error) {
	// Decode path: chatId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "chatId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateRoleParams is parameters of updateRole operation.
type UpdateRoleParams struct {
	// Role id.
	RoleId RoleId
	// Chat id.
	ChatId ChatId
}

func unpackUpdateRoleParams(packed middleware.Parameters) (params UpdateRoleParams) {
	{
		key := middleware.ParameterKey{
			Name: "roleId",
			In:   "path",
		}
		params.RoleId = packed[key].(RoleId)
	}
	{
		key := middleware.ParameterKey{
			Name: "chatId",
			In:   "query",
		}
		params.ChatId = packed[key].(ChatId)
	}
	return params
}

func decodeUpdateRoleParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateRoleParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: roleId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "roleId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotRoleIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotRoleIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.RoleId = RoleId(paramsDotRoleIdVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "roleId",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: chatId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "chatId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotChatIdVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotChatIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ChatId = ChatId(paramsDotChatIdVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "chatId",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
