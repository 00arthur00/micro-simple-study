package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcd/v2"
	yypservice "main.yapo.fun/service"
	"main.yapo.fun/serviceimpl"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	server := micro.NewService(
		micro.Name("api.yapo.fun.test"),
		micro.Registry(reg),
		micro.Address(":9091"),
	)
	yypservice.RegisterTestServiceHandler(server.Server(), new(serviceimpl.TestService))
	server.Run()
}
