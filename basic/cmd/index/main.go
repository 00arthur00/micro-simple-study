package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	router := gin.Default()
	router.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		data := make([]interface{}, 0)
		ctx.JSON(200, gin.H{
			"data": data,
		})
	})
	server := web.NewService(
		web.Address(":8000"),
		web.Handler(router),
	)
	server.Run()
}
