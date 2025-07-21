package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/task_manager/data"
	"github.com/yourusername/task_manager/models"
)

var jwtSecret = []byte("secret")

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := data.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func Login(c *gin.Context) {
	var creds models.User
	c.BindJSON(&creds)
	user, err := data.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func PromoteUser(c *gin.Context) {
	username := c.Param("username")
	if err := data.PromoteUser(username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user promoted"})
}

func CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	if err := data.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	tasks, _ := data.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	task, err := data.GetTaskByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var updated models.Task
	c.BindJSON(&updated)
	if err := data.UpdateTask(c.Param("id"), updated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func DeleteTask(c *gin.Context) {
	if err := data.DeleteTask(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
