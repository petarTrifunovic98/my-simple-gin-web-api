package controller

import (
	"fmt"
	"my-simple-gin-web-api/entities"
	"my-simple-gin-web-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControlller interface {
	Save(ctx *gin.Context) error
	FindAll() []entities.User
	FindByID(ctx *gin.Context) entities.User
}

type userController struct {
	service services.UserService
}

func New(service services.UserService) UserControlller {
	return &userController{
		service: service,
	}
}

func (c *userController) Save(ctx *gin.Context) error {
	var user entities.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}

func (c *userController) FindAll() []entities.User {
	return c.service.FindAll()
}

func (c *userController) FindByID(ctx *gin.Context) entities.User {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	fmt.Println("ID:", id)
	return c.service.FindByID(id)
}
