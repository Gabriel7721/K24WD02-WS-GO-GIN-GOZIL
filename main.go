package main

import (
	"fmt"
	"ws/src/common"

	"github.com/gin-gonic/gin"
)

func main() {
	common.LoadEnv()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to my chat server version 0.0.0.1")
	})

	port := common.GetEnv("PORT")
	fmt.Println("Server is running at http://localhost" + port)
	r.Run(port)
}
