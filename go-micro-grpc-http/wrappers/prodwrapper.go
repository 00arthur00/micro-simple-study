package wrappers

import (
	"context"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"services.yapo.fun/service"
)

type ProdWrapper struct {
	client.Client
}

func (w *ProdWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口")
	cfg := hystrix.CommandConfig{
		Timeout:                1000,
		RequestVolumeThreshold: 5,
		SleepWindow:            5000,
		ErrorPercentThreshold:  50,
	}
	cmdName := req.Service() + "." + req.Endpoint()
	hystrix.ConfigureCommand(cmdName, cfg)
	hystrix.Do(cmdName, func() error {
		return w.Client.Call(ctx, req, rsp, opts...)
	}, func(error) error {
		defaultList(rsp)
		return nil
	})
	return nil
}

func NewProdWrapper(c client.Client) client.Client {
	return &ProdWrapper{Client: c}
}

func defaultList(rsp interface{}) {
	switch t := rsp.(type) {
	case *service.ProdListResponse:
		list := make([]*service.ProdModel, 10, 10)
		for i := int32(0); i < 10; i++ {
			list[i] = &service.ProdModel{
				ProdID:   i,
				ProdName: fmt.Sprintf("prod%d", i),
			}
		}
		t.Data = list
	}

}
