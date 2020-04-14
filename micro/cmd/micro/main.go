package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
	yypservice "main.yapo.fun/service"
	"main.yapo.fun/serviceimpl"
)

func main() {
	reg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	server := grpc.NewService(
		service.Name("api.yapo.fun.test"),
		service.Registry(reg),
		service.Address(":9090"),
	)
	yypservice.RegisterTestServiceHandler(server.Server(), new(serviceimpl.TestService))
	server.Run()
}
