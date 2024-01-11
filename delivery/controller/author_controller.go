package controller

import (
	"net/http"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC usecase.AuthorUsecase
	rg       *gin.RouterGroup
}

func NewAuthorHandler(authorUC usecase.AuthorUsecase, rg *gin.RouterGroup) *AuthorController {
	return &AuthorController{
		authorUC: authorUC,
		rg:       rg,
	}
}

func (a *AuthorController) Route() {
	a.rg.GET("/authors/list/:id", a.ListAuthor)
	a.rg.GET("/authors/:id", a.GetAuthor)
}

func (a *AuthorController) ListAuthor(c *gin.Context) {
	id := c.Param("id")
	authors, err := a.authorUC.FindAllAuthor(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get author" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    authors,
	})
}

func (a *AuthorController) GetAuthor(c *gin.Context) {
	// if role admin muncul semua
	// if role bukan admin,muncul task nya sendiri
	id := c.Param("id")
	author, err := a.authorUC.FindAuthorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "author not found " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    author,
	})
}
