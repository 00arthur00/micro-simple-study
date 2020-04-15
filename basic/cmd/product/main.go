package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/v2"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcd/v2"
)

func makeIntArray(n int) []int {
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = i
	}
	return array
}
func main() {
	consulReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.Handle(http.MethodPost, "/prod", func(ctx *gin.Context) {
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
		web.Metadata(map[string]string{
			"protocol": "http",
		}),
	)
	server.Init()
	server.Run()
}
