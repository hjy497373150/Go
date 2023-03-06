package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message1" : "hello",
			"message2" : "gin",
		})
	})
	r.Run()
}