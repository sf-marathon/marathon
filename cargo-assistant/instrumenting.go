package cargo_assistant

import (
	"github.com/go-kit/kit/metrics"
	"time"
	"context"
	svc "marathon/cargo-assistant/service"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	svc.IGroupService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s svc.IGroupService) svc.IGroupService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		IGroupService:  s,
	}
}

func (s *instrumentingService) Group(ctx context.Context,id string) (error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Group").Add(1)
		s.requestLatency.With("method", "Group").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.IGroupService.Group(ctx,id)
}
func (s *instrumentingService) GetGroup(ctx context.Context) (*svc.GroupInfo, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetOrder").Add(1)
		s.requestLatency.With("method", "GetOrder").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.IGroupService.GetGroup(ctx)
}