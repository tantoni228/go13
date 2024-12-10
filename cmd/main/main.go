package main

import (
	"fmt"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/ogenregex"
	"github.com/ogen-go/ogen/validate"
)


type UserId string

func (s UserId) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes UserId from json.
func (s *UserId) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserId to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UserId(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s UserId) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserId) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}


var regexMap = map[string]ogenregex.Regexp{
	"uuid": ogenregex.MustCompile("uuid"),
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
        Regex:        regexMap["uuid"],
    }).Validate(string(alias)); err != nil {
        return errors.Wrap(err, "string")
    }
    return nil
}

func GenerateRandomUUID() (string) {
	id := uuid.New()
    return id.String()
}

func main() {
    uuid := UserId(GenerateRandomUUID())
    fmt.Println(uuid.Validate().Error())
}