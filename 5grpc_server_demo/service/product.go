package service

import "context"

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	//	实现业务逻辑

	return &ProductResponse{ProdStock: request.ProdId}, nil
}
