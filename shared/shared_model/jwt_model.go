package sharedmodel

import "github.com/golang-jwt/jwt/v5"

type CustomeClaims struct {
	jwt.RegisteredClaims
	AuthorID string
	Role     string
}
