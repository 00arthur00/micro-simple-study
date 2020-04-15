package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/client/http/v2"
	"github.com/micro/go-plugins/registry/etcd/v2"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	sel := selector.NewSelector(
		selector.SetStrategy(selector.Random),
		selector.Registry(reg),
	)
	//http client需要metadata为map[protocol:http]
	httpclient := http.NewClient(client.Selector(sel), client.ContentType("application/json"))
	//1.创建请求
	req := httpclient.NewRequest("httpprodservice", "/v1/prods", map[string]string{})
	//2.创建返回值
	rsp := make(map[string]interface{})
	if err := httpclient.Call(context.TODO(), req, &rsp); err != nil {
		log.Fatal(err)
	}
	//3.打印返回值
	fmt.Println(rsp)
}
