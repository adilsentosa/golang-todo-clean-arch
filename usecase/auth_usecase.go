package usecase

import (
	"todo-clean-arch/model/dto"
	"todo-clean-arch/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDTO) (dto.AuthResponseDTO, error)
}

type authUseCase struct {
	authorUC   AuthorUsecase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequestDTO) (dto.AuthResponseDTO, error) {
	author, err := a.authorUC.FindAuthorByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}

	if author.Password != payload.Password {
		return dto.AuthResponseDTO{}, err
	}

	// TODO generate jwt
	tokenDto, err := a.jwtService.GenerateToken(author)
	return dto.AuthResponseDTO{Token: tokenDto.Token}, nil
}

func NewAuthUseCase(authorUC AuthorUsecase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{
		authorUC:   authorUC,
		jwtService: jwtService,
	}
}
