package controller

import (
	"net/http"
	"user-system/pkg/dto"
	"user-system/pkg/middleware"
	"user-system/pkg/model"
	"user-system/pkg/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	UserService service.UserService
}

func (con *userController) AddBarReview(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	NewUser, err := con.UserService.AddUser(&user)
	if err != nil {
		middleware.HandleError(c, err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New User created with Title: " + NewUser.Email})

}

func (con *userController) DeleteUserById(c *gin.Context) {
	uid := c.Param("uid")

	err := con.UserService.DeleteUserById(uid)
	if err != nil {
		middleware.HandleError(c, err)
	}

	c.String(http.StatusOK, "User was deleted, id : %d", uid)
}

func (con *userController) UpdateUser(c *gin.Context) {
	uid := c.Param("uid")
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NewUser, err := con.UserService.UpdateUser(uid, &req)
	if err != nil {
		middleware.HandleError(c, err)
	}
	c.String(http.StatusOK, "User updated, id : %d", NewUser.ID)
}
