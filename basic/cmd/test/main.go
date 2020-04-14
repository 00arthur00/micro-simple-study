package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/00arthur00/micro/api/prod"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

func callAPI2(s selector.Selector) {
	myclient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myclient.NewRequest("prodservice", "/v1/prod", prod.ProdsRequest{Size: 1})
	var resp prod.ProdListResponse
	if err := myclient.Call(context.TODO(), req, &resp); err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetData())
}
func callAPI(addr string, path string, method string) (string, error) {
	req, err := http.NewRequest(method, "http://"+addr+path, nil)
	if err != nil {
		return "", err
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	myselector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	for {
		callAPI2(myselector)
		time.Sleep(time.Second)
	}

}
func main1() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	for {
		services, err := consulReg.GetService("prodservice")
		if err != nil {
			log.Fatal(err)
		}
		next := selector.RoundRobin(services)
		node, err := next()
		if err != nil {
			log.Println(err)
		}

		fmt.Println(node.Address, node.Id, node.Metadata)
		fmt.Println(callAPI(node.Address, "/v1/prod", "GET"))
		time.Sleep(1 * time.Second)
	}
}
