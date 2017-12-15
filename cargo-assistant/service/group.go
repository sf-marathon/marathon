package service

import (
	"context"
	"fmt"
	"marathon/cargo-assistant/dao"
)

type GroupInfo struct {
	Group      *dao.Group
	ProMktBase *dao.ProMarketBase
}

type IGroupService interface {
	Group(ctx context.Context, proMktBaseId string) error
	GetGroup(ctx context.Context) (*GroupInfo, error)
}

type GroupService struct {
	groupDao      dao.IGroupDao
	proMktBaseDao dao.IProMarketBaseDao
}

func NewGroupService(groupDao dao.IGroupDao, proMktBaseDao dao.IProMarketBaseDao) *GroupService {
	return &GroupService{
		groupDao:      groupDao,
		proMktBaseDao: proMktBaseDao,
	}
}

func (s *GroupService) GetGroup(ctx context.Context) (*GroupInfo, error) {
	group, errg := s.groupDao.Select()
	pmb, errp := s.proMktBaseDao.Select(group.MarketId)
	var err error
	if errg != nil || errp != nil {
		err = fmt.Errorf("%v %v", errg, errp)
	}
	return &GroupInfo{Group:group, ProMktBase:pmb}, err
}

func (s *GroupService) Group(ctx context.Context, proMktBaseId string) error {
	pmb, err := s.proMktBaseDao.Select(proMktBaseId)
	if err != nil {
		return err
	}
	return s.groupDao.Insert(pmb)
}
