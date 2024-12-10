// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

func encodeChangePasswordResponse(response ChangePasswordRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *ChangePasswordNoContent:
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

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCheckTokenResponse(response CheckTokenRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *CheckTokenNoContent:
		w.WriteHeader(204)

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

func encodeGetMeResponse(response GetMeRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *User:
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

func encodeGetUserByIdResponse(response GetUserByIdRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *User:
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

	case *UserNotFoundResponse:
		w.WriteHeader(404)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeSignInResponse(response SignInRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *SignInResponse:
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

	case *SignInUnauthorized:
		w.WriteHeader(401)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeSignUpResponse(response SignUpRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *SignUpNoContent:
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

	case *SignUpConflict:
		w.WriteHeader(409)

		return nil

	case *InternalErrorResponse:
		w.WriteHeader(500)

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdateMeResponse(response UpdateMeRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *User:
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
