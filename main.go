package main

import (
	"fmt"
	"net/http"

	"github.com/fickkmdev/backend-golang101/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//middleware ต้องทำในmiddlewareก่อน ถึงจะทำโค้ดทุกบรรทัดหลังจากนี้
	r.Use(middlewareFunc())
	{
		r.GET("/welcome", welcome)
		r.POST("/welcome-message", welcomeMessage)
		r.POST("/welcome-message-json", welcomeMessageJson)
	}

	r.Run(":3030")
}

func middlewareFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//middleware
		fmt.Println("middleware doing...")
	}
}

func welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello, world",
	})
}

func welcomeMessage(ctx *gin.Context) {
	var result model.Result
	result.ID = ctx.Query("id")
	result.Name = ctx.PostForm("name")
	result.Message = ctx.PostForm("message")
	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func welcomeMessageJson(ctx *gin.Context) {
	var input model.Result
	e := ctx.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"id":      input.ID,
			"name":    input.Name,
			"message": input.Message,
		},
	)
}
