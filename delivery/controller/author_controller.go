package controller

import (
	"fmt"
	"net/http"
	"todo-clean-arch/model"
	"todo-clean-arch/shared/common"
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
	a.rg.PUT("/authors/:id", a.UpdateAuthor)
	a.rg.DELETE("/authors/:id")
}

func (a *AuthorController) RemoveAuthor(c *gin.Context) {
	id := c.Param("id")
	err := a.authorUC.RemoveAuthor(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("failed to delete %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "succes",
	})
}

func (a *AuthorController) UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author model.Author
	err := c.ShouldBind(&author)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid json %v", err))
		return
	}

	if author.Name == "" || author.Email == "" || author.Password == "" || author.Role == "" {
		common.SendErrorResponse(c, http.StatusBadRequest, "all field cant be empty")
		return
	}

	author.ID = id
	resAuthor, err := a.authorUC.UpdateAuthor(author)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to update %v ", err))
		return
	}
	common.SendSingleResponse(c, resAuthor, "Success")
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
