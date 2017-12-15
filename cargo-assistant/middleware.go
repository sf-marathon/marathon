package cargo_assistant

import (
	"context"
	_ "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	svc "marathon/cargo-assistant/service"
)

type LoggingMiddleware struct {
	next   svc.IGroupService
	logger log.Logger
}

func NewLoggingMiddleware(logger log.Logger, s svc.IGroupService) *LoggingMiddleware {
	return &LoggingMiddleware{
		next:   s,
		logger: logger,
	}
}

func (l *LoggingMiddleware) Group(ctx context.Context, id string) error {
	//TODO log
	l.logger.Log("Group",id)
	return l.next.Group(ctx, id)
}

func (l *LoggingMiddleware) GetGroup(ctx context.Context, id string) (*svc.GroupInfo, error) {
	//TODO log
	l.logger.Log("GetGroup",id)
	return l.next.GetGroup(ctx)
}
