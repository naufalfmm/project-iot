package jwt

import (
	"github.com/dgrijalva/jwt-go"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/naufalfmm/project-iot/common/consts"
)

func (j *jwtHelper) Parse(tokenString string, claims jwtGo.Claims) (jwtGo.Claims, error) {
	token, err := jwtGo.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.verifyKey, nil
	})
	if err != nil || token == nil || token.Claims == nil {
		return nil, consts.UnclaimedToken
	}

	claimsRes := token.Claims

	return claimsRes, nil
}
