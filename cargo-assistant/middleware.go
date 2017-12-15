package cargo_assistant

import (
	"context"
	_ "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	next   IOrderService
	logger log.Logger
}

func NewLoggingMiddleware(logger log.Logger, s IOrderService) *LoggingMiddleware {
	return &LoggingMiddleware{
		next:   s,
		logger: logger,
	}
}

func (l *LoggingMiddleware) Order(ctx context.Context, r *Order) error {
	//TODO log
	l.logger.Log("order",r)
	return l.next.Order(ctx, r)
}

func (l *LoggingMiddleware) GetOrder(ctx context.Context, id string) (*Order, error) {
	//TODO log
	return l.next.GetOrder(ctx, id)
}
