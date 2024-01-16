package service

import (
	"time"
	"todo-clean-arch/config"
	"todo-clean-arch/model"
	"todo-clean-arch/model/dto"
	sharedmodel "todo-clean-arch/shared/shared_model"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(author model.Author) (dto.AuthResponseDTO, error)
}

type jwtService struct {
	cfg config.TokenConfig
}

func (j *jwtService) GenerateToken(author model.Author) (dto.AuthResponseDTO, error) {
	claims := sharedmodel.CustomeClaims{
		AuthorID: author.ID,
		Role:     author.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtExpiresTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}
	return dto.AuthResponseDTO{
		Token: tokenString,
	}, nil
}

func NewJwtService() JwtService {
	return &jwtService{}
}
