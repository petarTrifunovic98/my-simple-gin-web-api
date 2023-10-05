package main

import (
	"fmt"
	"io"
	"my-simple-gin-web-api/controller"
	"my-simple-gin-web-api/dbdriver"
	"my-simple-gin-web-api/repository"
	"my-simple-gin-web-api/services"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	err            error
	connection     net.Conn
	client         *dbdriver.MySimpleDatabaseClient
	userRepository repository.UserRepository
	userService    services.UserService
	userController controller.UserControlller
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.Default()

	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	connection, err = net.Dial("tcp", "localhost:9988")
	if err != nil {
		fmt.Printf("An error in tcp dial occured %v\n", err)
		os.Exit(1)
	}

	client = dbdriver.NewSimpleDatabaseClient(connection)
	userRepository = repository.NewUserRepository(client)
	userService = services.New(userRepository)
	userController = controller.New(userService)

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindByID(ctx))
	})

	server.POST("/users", func(ctx *gin.Context) {
		err := userController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User input is valid"})
		}
	})

	server.Run(":8080")
}
