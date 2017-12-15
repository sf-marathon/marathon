package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	kitlog "github.com/go-kit/kit/log"
	"strconv"
	"time"
)

const TABLE_NAME_GROUP = "group"

type Group struct {
	GroupId    int       `json:"group_id" orm:"column(group_id);pk"`
	MarketId   string    `json:"market_id" orm:"column(market_id)"`
	GroupName  string    `json:"group_name" orm:"column(group_name)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	DueTime    time.Time `json:"due_time" orm:"column(due_time)"`
}
type IGroupDao interface {
	Select(id string) (*Group, error)
	Insert(proMktBase *ProMarketBase) error
}
type GroupDao struct {
	log kitlog.Logger
}

func NewGroupDao(log kitlog.Logger) (*GroupDao, error) {
	groupDao := &GroupDao{
		log: log,
	}
	return groupDao, nil
}

func (m *GroupDao) Insert(proMktBase *ProMarketBase) error {
	o := orm.NewOrm()
	// find last group of this base
	//TODO: finish this
	_, err := o.Insert(&Group{})
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot insert record err: %v", err))
		return err
	}
	return nil
}

func (m *GroupDao) Select(id string) (*Group, error) {
	group := Group{}
	intId, err := strconv.Atoi(id)
	if err != nil {
		m.log.Log(fmt.Sprintf("Error parsing group id %s, err: %v", id, err))
		return &group, err
	}
	group.GroupId = intId
	o := orm.NewOrm()
	err = o.QueryTable(TABLE_NAME_GROUP).One(&group)
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot find record id %s, err: %v", id, err))
		return &group, err
	}
	return &group, nil
}
