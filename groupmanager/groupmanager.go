package main

import (
	kitlog "github.com/go-kit/kit/log"

	"context"
	"marathon/cargo-assistant/dao"
	"time"
)

type GroupManager struct {
	groupDao      dao.IGroupDao
	joinDao       dao.IJoinDao
	proMktBaseDao dao.IProMarketBaseDao
	checkInterval time.Duration
	logger        kitlog.Logger
}

func NewGroupManager(groupDao dao.IGroupDao, joinDao dao.IJoinDao,
	proMktBaseDao dao.IProMarketBaseDao,
	inv time.Duration, logger kitlog.Logger) *GroupManager {
	return &GroupManager{
		groupDao:      groupDao,
		joinDao:       joinDao,
		proMktBaseDao: proMktBaseDao,
		checkInterval: inv,
		logger:        logger,
	}
}

func (m *GroupManager) Run(ctx context.Context) {
	m.logger.Log("GroupManager start to run.")
	m.cycleGroup(ctx)
	m.logger.Log("GroupManager stopped.")
}

func (m *GroupManager) cycleGroup(ctx context.Context) {
loop:
	for {
		nextDue := time.Now().Add(m.checkInterval)
		grp, err := m.groupDao.Select()
		if err != nil {
			grp = nil
			m.logger.Log("Error get current group")
		} else {
			nextDue = grp.DueTime
		}
		m.logger.Log("nextdue", nextDue)
		interval := nextDue.Sub(time.Now())
		m.logger.Log("sleep for", interval)
		select {
		case <-ctx.Done():
			break loop
		case <-time.After(interval):
		}
		m.logger.Log("due for", grp)
		m.nextGroup()
	}
}

func (m *GroupManager) nextGroup() {
	grp, err := m.groupDao.Select()
	if err != nil {
		m.logger.Log("Error get current group while while nextGroup()")
		return
	}
	mkt, err := m.proMktBaseDao.Select(grp.MarketId)
	if err != nil {
		m.logger.Log("Error get current market while while nextGroup()")
		return
	}
	err = m.groupDao.Insert(mkt)
	if err != nil {
		m.logger.Log("Error create current market while nextGroup()")
		return
	}
	return
}
