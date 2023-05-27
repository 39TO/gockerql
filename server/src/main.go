package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	doMain()
}

func doMain() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	r.Run(":8082")
}
