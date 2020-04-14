package serviceimpl

import (
	"context"
	"fmt"

	"main.yapo.fun/service"
)

type TestService struct {
}

func (t *TestService) Call(ctx context.Context, req *service.Request, rsp *service.Response) error {
	fmt.Println("调用函数")
	rsp.Data = "hello"
	return nil
}
