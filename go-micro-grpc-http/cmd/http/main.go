package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcd/v2"
	"services.yapo.fun/service"
	"services.yapo.fun/wrappers"
)

func defaultList() *service.ProdListResponse {
	list := make([]*service.ProdModel, 10, 10)
	for i := int32(0); i < 10; i++ {
		list[i] = &service.ProdModel{
			ProdID:   i,
			ProdName: fmt.Sprintf("prod%d", i),
		}
	}
	return &service.ProdListResponse{Data: list}
}
func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	sel := selector.NewSelector(
		selector.SetStrategy(selector.RoundRobin),
		selector.Registry(etcdReg),
	)
	myservice := micro.NewService(
		micro.Name("prodservice.client"),
		// micro.Registry(etcdReg),
		micro.WrapClient(NewLogWrapper, wrappers.NewProdWrapper),
		micro.Selector(sel),
	)
	prodservice := service.NewProdService(
		"prodservice",
		myservice.Client(),
	)

	router := gin.Default()
	g := router.Group("/v1")
	g.Handle(http.MethodPost, "/prods", func(ginctx *gin.Context) {
		var req service.ProdRequest
		if err := ginctx.Bind(&req); err != nil {
			ginctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		resp, err := prodservice.GetProdsList(context.TODO(), &req)
		// //熔断代码改造
		// //1. 配置config
		// config := hystrix.CommandConfig{
		// 	Timeout: 1000, //毫秒
		// }
		// //2.配置command
		// hystrix.ConfigureCommand("getprods", config)
		// //3. do方法(同步)
		// var resp *service.ProdListResponse
		// err := hystrix.Do("getprods", func() error {
		// 	var err error
		// 	resp, err = prodservice.GetProdsList(context.TODO(), &req)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	return nil
		// }, func(err error) error {
		// 	resp = defaultList()
		// 	return nil
		// })
		if err != nil {
			ginctx.AbortWithError(http.StatusInternalServerError, err)
			ginctx.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
			})
			return
		}
		ginctx.JSON(http.StatusOK, resp)
	})

	server := web.NewService(
		web.Name("httpprodservice"),
		web.Address(":8081"),
		web.Handler(router),
		web.Registry(etcdReg),
		//需要设置，否则，micro http client无法连接
		web.Metadata(map[string]string{
			"protocol": "http",
		}),
	)
	server.Init()
	server.Run()
}
