package usecase

import "todo-clean-arch/model/dto"

type AuthUseCase interface {
	Login(payload dto.AuthRequestDTO) (dto.AuthResponseDTO, error)
}

type authUseCase struct {
	authorUC AuthorUsecase
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
	return dto.AuthResponseDTO{Token: ""}, nil
}

func NewAuthUseCase(authorUC AuthorUsecase) AuthUseCase {
	return &authUseCase{
		authorUC: authorUC,
	}
}
