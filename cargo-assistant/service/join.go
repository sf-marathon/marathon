package service

import (
	"context"
	"marathon/cargo-assistant/dao"
)

type IJoinService interface {
	Join(ctx context.Context, join *dao.Join) error
	GetJoin(ctx context.Context, groupId string) ([]*dao.Join, error)
}

type JoinService struct {
	joinDao dao.IJoinDao
}

func NewJoinService(joinDao dao.IJoinDao) *JoinService {
	return &JoinService{
		joinDao: joinDao,
	}
}

func (s *JoinService) GetJoin(ctx context.Context, groupId string) ([]*dao.Join, error) {
	return s.joinDao.Select(groupId)
}

func (s *JoinService) Join(ctx context.Context, join *dao.Join) error {
	return s.joinDao.Insert(join)
}
