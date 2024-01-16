package controller

import (
	"net/http"
	"todo-clean-arch/model/dto"
	"todo-clean-arch/shared/common"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload dto.AuthRequestDTO
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// send payload to usecase
	common.SendSingleResponse(c, nil, "login success")
}

func (a *AuthController) Route() {
	a.rg.POST("/login", a.loginHandler)
}

func NewAuthController(authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		authUC: authUC,
		rg:     rg,
	}
}
