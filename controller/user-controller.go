package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Save(user)
	return user
}
