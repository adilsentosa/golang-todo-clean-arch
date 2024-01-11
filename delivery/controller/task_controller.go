package controller

import (
	"fmt"
	"net/http"
	"todo-clean-arch/model"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskUC usecase.TaskUsecase
}

func NewTaskHandler(taskUC usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUC: taskUC,
	}
}

func (t *TaskHandler) Route() {
}

func (a *TaskHandler) CreateHandler(c *gin.Context) {
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

func (a *TaskHandler) FindTaskByAuthor(c *gin.Context) {
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
