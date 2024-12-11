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

// GetUserByIdParams is parameters of getUserById operation.
type GetUserByIdParams struct {
	// User id.
	UserId UserId
}

func unpackGetUserByIdParams(packed middleware.Parameters) (params GetUserByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(UserId)
	}
	return params
}

func decodeGetUserByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserByIdParams, _ error) {
	// Decode path: userId.
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
