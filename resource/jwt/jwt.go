package jwt

import (
	"crypto/rsa"
	"io/ioutil"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/naufalfmm/project-iot/resource/config"
)

type (
	JWT interface {
		Parse(tokenString string, claims jwtGo.Claims) (jwtGo.Claims, error)
		Create(claims jwtGo.Claims) (string, error)
	}

	jwtHelper struct {
		verifyKey *rsa.PublicKey
		signKey   *rsa.PrivateKey
	}
)

func New(config *config.EnvConfig) (JWT, error) {
	signBytes, err := ioutil.ReadFile(config.JwtPrivateKey)
	if err != nil {
		return nil, err
	}

	signKey, err := jwtGo.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}

	verifyBytes, err := ioutil.ReadFile(config.JwtPublicKey)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwtGo.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return &jwtHelper{
		verifyKey: verifyKey,
		signKey:   signKey,
	}, nil
}
