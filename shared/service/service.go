package service

import (
	"time"
	"todo-clean-arch/model"
	"todo-clean-arch/model/dto"
	sharedmodel "todo-clean-arch/shared/shared_model"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(author model.Author) (dto.AuthResponseDTO, error)
}

type jwtService struct{}

func (j *jwtService) GenerateToken(author model.Author) (dto.AuthResponseDTO, error) {
	claims := sharedmodel.CustomeClaims{
		AuthorID: author.ID,
		Role:     author.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "enigmacamp.com",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Hour * 2),
		},
	}
	return dto.AuthResponseDTO{}, nil
}

func NewJwtService() JwtService {
	return &jwtService{}
}
