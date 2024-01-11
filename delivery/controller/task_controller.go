package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-clean-arch/model"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUC usecase.TaskUsecase
	rg     *gin.RouterGroup
}

func NewTaskHandler(taskUC usecase.TaskUsecase, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		taskUC: taskUC,
		rg:     rg,
	}
}

func (t *TaskController) Route() {
	t.rg.GET("/tasks/list", t.ListHandler)
	t.rg.GET("/tasks/get/:id", t.FindTaskByAuthor)
	t.rg.POST("/tasks/create", t.CreateHandler)
}

func (a *TaskController) CreateHandler(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Errorf("failed to bind json %v", err),
		})
		return
	}

	if task.AuthorID == "" || task.Title == "" || task.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "AuthorID or title or content is empty",
		})
		return
	}

	task, err := a.taskUC.RegisterNewTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "failed to create task %v " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "task created successfuly",
		"data":    task,
	})
}

func (a *TaskController) FindTaskByAuthor(c *gin.Context) {
	author := c.Param("id")
	tasks, err := a.taskUC.FindTaskByAuthor(author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "task with authorID: " + author + " not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    tasks,
	})
}

func (t *TaskController) ListHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	tasks, paging, err := t.taskUC.FindAllTask(page, size)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusOK,
			"message": "Ok",
		},
		"data":   tasks,
		"paging": paging,
	})
}
