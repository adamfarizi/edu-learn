package modelutils

import "github.com/golang-jwt/jwt/v5"

type JWTPayloadClaim struct {
	jwt.RegisteredClaims
	UserId int
	Role   string
}
