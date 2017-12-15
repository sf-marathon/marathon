package service

import (
	"context"
	"marathon/cargo-assistant/dao"
)

type IOrderService interface {
	Order(ctx context.Context, r *dao.ProMarketBase) error
	GetOrder(ctx context.Context, id string) (*dao.ProMarketBase, error)
}

/*type OrderService struct {
	orderDao IOrderDao
}

func NewOrderService(orderDao IOrderDao) *OrderService {
	return &OrderService{
		orderDao: orderDao,
	}
}



func (rs *OrderService) GetProMktBase(ctx context.Context, id string) (*dao.ProMarketBase, error) {
	return rs.orderDao.Select(id)
}*/
