package main

import (
	"fmt"

	"ws/src/auth"
	"ws/src/common"
	"ws/src/user"

	"github.com/gin-gonic/gin"
)

func main() {
	common.LoadEnv()
	userRepo := user.NewRepository(common.MongoConnect())
	userController := user.NewController(userRepo)
	authController := auth.NewController(userRepo)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to my chat server version 0.0.0.1")
	})
	r.POST("/api/register", userController.Register)
	r.POST("/api/login", authController.Login)

	port := common.GetEnv("PORT")
	fmt.Println("Server is running at http://localhost" + port)
	r.Run(port)
}
