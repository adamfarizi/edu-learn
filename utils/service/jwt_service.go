package service

import (
	"fmt"
	"edu-learn/config"
	"edu-learn/model"
	modelutils "edu-learn/utils/model_utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	cfg config.TokenConfig
}

type JwtService interface {
	CreateToken(user model.User) (string, error)
	VerifyToken(tokenString string) (modelutils.JWTPayloadClaim, error)
}

func (j *jwtService) CreateToken(user model.User) (string, error) {
	tokenKey := j.cfg.JWTSignatureKey

	claims := modelutils.JWTPayloadClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.cfg.ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.AccessTokenLifeTime)),
		},
		UserId: user.ID,
		Role:   user.Role,
	}

	jwtNewClaim := jwt.NewWithClaims(j.cfg.JWTSigningMethod, claims)

	token, err := jwtNewClaim.SignedString(tokenKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) VerifyToken(tokenString string) (modelutils.JWTPayloadClaim, error) {
	tokenParse, err := jwt.ParseWithClaims(tokenString, &modelutils.JWTPayloadClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return j.cfg.JWTSignatureKey, nil
		},
	)
	if err != nil {
		return modelutils.JWTPayloadClaim{}, err
	}

	claim, ok := tokenParse.Claims.(*modelutils.JWTPayloadClaim)
	if !ok {
		return modelutils.JWTPayloadClaim{}, fmt.Errorf("error claim")
	}

	return *claim, nil
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg}
}