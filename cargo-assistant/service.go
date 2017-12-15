package cargo_assistant

import (
	"context"
)

type IOrderService interface {
	Order(ctx context.Context, r *Order) error
	GetOrder(ctx context.Context, id string) (*Order, error)
}

type OrderService struct {
	orderDao IOrderDao
}

func NewOrderService(orderDao IOrderDao) *OrderService {
	return &OrderService{
		orderDao: orderDao,
	}
}

func (rs *OrderService) Order(ctx context.Context, r *Order) error {
	return rs.orderDao.Insert(r)
}

func (rs *OrderService) GetOrder(ctx context.Context, id string) (*Order, error) {
	return rs.orderDao.Select(id)
}
