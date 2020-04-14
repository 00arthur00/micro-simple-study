package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func makeIntArray(n int) []int {
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = i
	}
	return
}
func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.Handle(http.MethodPost, "/prod", func(ctx *gin.Context) {
			ctx.Bind()
			ctx.JSON(http.StatusOK, gin.H{
				"data": makeIntArray(10),
			})
		})
	}
	server := web.NewService(
		web.Name("prodservice"),
		// web.Address(":8001"),
		web.Handler(router),
		web.Registry(consulReg),
	)
	server.Init()
	server.Run()
}
