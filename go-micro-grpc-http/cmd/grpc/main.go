package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcd/v2"
	"services.yapo.fun/service"
	"services.yapo.fun/serviceimpl"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	webservice := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(reg),
	)
	webservice.Init()
	service.RegisterProdServiceHandler(webservice.Server(), new(serviceimpl.ProdService))
	webservice.Run()
}
