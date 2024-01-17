package usecase

import (
	"log"
	"todo-clean-arch/model/dto"
	"todo-clean-arch/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDTO) (dto.AuthResponseDTO, error)
	// GetKey() []byte
}

type authUseCase struct {
	authorUC   AuthorUsecase
	jwtService service.JwtService
}

// func (a *authUseCase) GetKey() []byte {
// 	return a.jwtService.GetKey()
// }

func (a *authUseCase) Login(payload dto.AuthRequestDTO) (dto.AuthResponseDTO, error) {
	// fmt.Println("authUC Login")
	// fmt.Println(payload)
	author, err := a.authorUC.FindAuthorByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}
	log.Println("authUC.FindAuthorByEmail", author)
	log.Println(payload.Password == author.Password)

	if payload.Password != author.Password {
		return dto.AuthResponseDTO{}, err
	}
	log.Println("password correct")

	// TODO generate jwt
	tokenDto, err := a.jwtService.GenerateToken(author)
	log.Println("TokenDto")
	log.Println(tokenDto)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}
	return tokenDto, nil
}

func NewAuthUseCase(authorUC AuthorUsecase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{
		authorUC:   authorUC,
		jwtService: jwtService,
	}
}
