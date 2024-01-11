package common

import (
	"net/http"
	sharedmodel "todo-clean-arch/shared/shared_model"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse() {
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, sharedmodel.SingleResponse{})
}
