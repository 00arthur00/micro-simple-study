package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"services.yapo.fun/service"
	"services.yapo.fun/serviceimpl"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	webservice := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
	)
	webservice.Init()
	service.RegisterProdServiceHandler(webservice.Server(), new(serviceimpl.ProdService))
	webservice.Run()
}
