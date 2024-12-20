// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *ChangePasswordReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.OldPassword.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "old_password",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.NewPassword.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "new_password",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s Email) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        true,
		Hostname:     false,
		Regex:        regexMap["^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s Password) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        regexMap["^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,25}$"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s *SignInReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Password.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *SignUpReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Password.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *User) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.ID.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "id",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s UserId) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    0,
		MinLengthSet: false,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        regexMap["^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"],
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s *UserInput) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    20,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Username)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "username",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
