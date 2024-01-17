package middleware_mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthorMiddlewareMock struct {
	mock.Mock
}

func (a *AuthorMiddlewareMock) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
