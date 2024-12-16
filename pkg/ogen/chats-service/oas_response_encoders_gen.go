// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

func encodeBanUserResponse(response BanUserRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *BanUserNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *BanUserNotFound:
		w.WriteHeader(404)

		return nil

	case *BanUserConflict:
		w.WriteHeader(409)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCheckAccessResponse(response CheckAccessRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *CheckAccessNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *CheckAccessNotFound:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreateChatResponse(response CreateChatRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Chat:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreateRoleResponse(response CreateRoleRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Role:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *CreateRoleConflict:
		w.WriteHeader(409)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteChatResponse(response DeleteChatRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *DeleteChatNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteRoleResponse(response DeleteRoleRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *DeleteRoleNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *DeleteRoleNotFound:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetChatByIdResponse(response GetChatByIdRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Chat:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetJoinCodeResponse(response GetJoinCodeRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *JoinCodeResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetMyRoleResponse(response GetMyRoleRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Role:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetRoleByIdResponse(response GetRoleByIdRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Role:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *GetRoleByIdNotFound:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeJoinChatResponse(response JoinChatRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *JoinChatNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *JoinChatConflict:
		w.WriteHeader(409)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeLeaveChatResponse(response LeaveChatRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *LeaveChatNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListBannedUsersResponse(response ListBannedUsersRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *ListBannedUsersOKApplicationJSON:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListChatsResponse(response ListChatsRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *ListChatsOKApplicationJSON:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListMembersResponse(response ListMembersRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *ListMembersOKApplicationJSON:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListRolesResponse(response ListRolesRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *ListRolesOKApplicationJSON:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeSetRoleResponse(response SetRoleRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *SetRoleNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *SetRoleNotFound:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUnbanUserResponse(response UnbanUserRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *UnbanUserNoContent:
		w.WriteHeader(204)

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *UnbanUserNotFound:
		w.WriteHeader(404)

		return nil

	case *UnbanUserConflict:
		w.WriteHeader(409)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdateChatResponse(response UpdateChatRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Chat:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *ChatNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdateRoleResponse(response UpdateRoleRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Role:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InvalidInputResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthenticatedResponse:
		w.WriteHeader(401)

		return nil

	case *UnauthorizedResponse:
		w.WriteHeader(403)

		return nil

	case *UpdateRoleNotFound:
		w.WriteHeader(404)

		return nil

	case *UpdateRoleConflict:
		w.WriteHeader(409)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
