package serviceimpl

import (
	"context"
	"fmt"
	"time"

	"services.yapo.fun/service"
)

type ProdService struct {
}

func (p *ProdService) GetProdsList(ctx context.Context, req *service.ProdRequest, resp *service.ProdListResponse) error {
	time.Sleep(3 * time.Second)
	if req.Size <= 0 {
		req.Size = 2
	}
	list := make([]*service.ProdModel, req.Size, req.Size)
	for i := int32(0); i < req.Size; i++ {
		list[i] = &service.ProdModel{
			ProdID:   i,
			ProdName: fmt.Sprintf("prod%d", i),
		}
	}
	resp.Data = list
	return nil
}
