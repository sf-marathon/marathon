package cargo_assistant

import (
	"github.com/go-kit/kit/metrics"
	"time"
	"context"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	IOrderService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s IOrderService) IOrderService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		IOrderService:  s,
	}
}

func (s *instrumentingService) Order(ctx context.Context,r *Order) (error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Order").Add(1)
		s.requestLatency.With("method", "Order").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.IOrderService.Order(ctx,r)
}
func (s *instrumentingService) GetOrder(ctx context.Context, id string) (*Order, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetOrder").Add(1)
		s.requestLatency.With("method", "GetOrder").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.IOrderService.GetOrder(ctx,id)
}