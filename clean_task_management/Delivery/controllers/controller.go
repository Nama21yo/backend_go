package controllers

import (
	"net/http"
	"task-manager/Usecases"
	"time"

	"github.com/gin-gonic/gin"
)

// UserController wires HTTP requests to user usecases
type UserController struct {
	UserUC Usecases.UserUsecase
	JWTSvc infrastructure.JWTService
}

type registerReq struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	DisplayName string `json:"display_name"`
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewUserController(uc Usecases.UserUsecase, jwt infrastructure.JWTService) *UserController {
	return &UserController{UserUC: uc, JWTSvc: jwt}
}

func (u *UserController) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.UserUC.Register(req.Email, req.Password, req.DisplayName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// hide sensitive fields
	user.PasswordHash = ""
	c.JSON(http.StatusCreated, user)
}

func (u *UserController) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := u.UserUC.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user.PasswordHash = ""
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
		"ttl":   time.Now().Add(time.Hour * 24).Format(time.RFC3339),
	})
}

// TaskController wires HTTP requests to task usecases
type TaskController struct {
	TaskUC Usecases.TaskUsecase
}

func NewTaskController(tu Usecases.TaskUsecase) *TaskController {
	return &TaskController{TaskUC: tu}
}

type createTaskReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func getCurrentUserID(c *gin.Context) string {
	v, _ := c.Get(infrastructure.ContextUserIDKey)
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var req createTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ownerID := getCurrentUserID(c)
	task, err := t.TaskUC.CreateTask(req.Title, req.Description, ownerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requester := getCurrentUserID(c)
	updated, err := t.TaskUC.UpdateTask(&task, requester)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (t *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	requester := getCurrentUserID(c)
	task, err := t.TaskUC.GetTaskByID(id, requester)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (t *TaskController) ListTasks(c *gin.Context) {
	requester := getCurrentUserID(c)
	tasks, err := t.TaskUC.ListTasks(requester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	requester := getCurrentUserID(c)
	if err := t.TaskUC.DeleteTask(id, requester); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
