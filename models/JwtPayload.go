package models

import (
	"github.com/gbrlsnchs/jwt"
)

type JwtPayload struct {
	jwt.Payload
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"bar,omitempty"`
}