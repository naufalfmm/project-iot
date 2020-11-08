package jwt

import (
	"github.com/dgrijalva/jwt-go"
	jwtGo "github.com/dgrijalva/jwt-go"
)

func (j *jwtHelper) Create(claims jwtGo.Claims) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = claims

	return t.SignedString(j.signKey)
}
