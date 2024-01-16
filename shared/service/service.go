package service

import (
	"todo-clean-arch/model"
	"todo-clean-arch/model/dto"
)

type JwtService interface {
	GenerateToken(author model.Author) (dto.AuthResponseDTO, error)
}

type jwtService struct{}

func (j *jwtService) GenerateToken(author model.Author) (dto.AuthResponseDTO, error) {
	panic("not implemented")
}

func NewJwtService() JwtService {
	return &jwtService{}
}
