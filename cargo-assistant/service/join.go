package service

import (
	"context"
	"marathon/cargo-assistant/dao"
	"fmt"
)

type IJoinService interface {
	Join(ctx context.Context, join *dao.Join) error
	GetJoin(ctx context.Context, groupId string) ([]*dao.Join, error)
}

type JoinService struct {
	joinDao dao.IJoinDao
	groupDao dao.IGroupDao
	proMktBaseDao dao.IProMarketBaseDao
}

func NewJoinService(joinDao dao.IJoinDao, groupDao dao.IGroupDao, baseDao dao.IProMarketBaseDao) *JoinService {
	return &JoinService{
		joinDao: joinDao,
		groupDao: groupDao,
		proMktBaseDao:baseDao,
	}
}

func (s *JoinService) GetJoin(ctx context.Context, groupId string) ([]*dao.Join, error) {
	return s.joinDao.Select(groupId)
}

func (s *JoinService) Join(ctx context.Context, join *dao.Join) error {
	err := s.joinDao.Insert(join)
	if err != nil {
		return err
	}
	//check num and insert new group

	group,err := s.groupDao.Select()
	if err != nil {
		return err
	}
	joins,err := s.joinDao.Select(fmt.Sprintf("%d",group.GroupId))
	if err != nil {
		return err
	}
	pmb,err := s.proMktBaseDao.Select(group.MarketId)
	fmt.Println("CURRENT:",len(joins),pmb.GroupLimit)
	if len(joins) >= pmb.GroupLimit {
		return s.groupDao.Insert(pmb)
	}
	return nil
}
