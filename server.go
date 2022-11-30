package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()
	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())

	})
	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	server.Run(":8080")
}
