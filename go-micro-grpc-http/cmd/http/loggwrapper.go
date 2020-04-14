package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/client"
)

type logWrapper struct {
	client.Client
}

func (lw *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口")
	return lw.Client.Call(ctx, req, rsp, opts...)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{Client: c}
}
