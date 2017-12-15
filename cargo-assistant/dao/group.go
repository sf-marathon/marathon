package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	kitlog "github.com/go-kit/kit/log"
	"time"
)

const (
	TABLE_NAME_GROUP = "group"

	COLUMN_DUE_TIME = "due_time"
	COLUMN_GROUP_ID = "group_id"
)

type Group struct {
	GroupId    int       `json:"group_id" orm:"column(group_id);pk;auto"`
	MarketId   string    `json:"market_id" orm:"column(market_id)"`
	GroupName  string    `json:"group_name" orm:"column(group_name)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	DueTime    time.Time `json:"due_time" orm:"column(due_time)"`
}
type IGroupDao interface {
	Select() (*Group, error)
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

func (m *GroupDao) Select() (*Group, error) {
	/*	group := Group{}
		intId, err := strconv.Atoi(id)
		if err != nil {
			m.log.Log(fmt.Sprintf("Error parsing group id %s, err: %v", id, err))
			return &group, err
		}
		fmt.Println(intId)
		group.GroupId = intId*/
	var groups []*Group
	o := orm.NewOrm()
	_, err := o.QueryTable(TABLE_NAME_GROUP).OrderBy("-" + COLUMN_GROUP_ID).Limit(1).All(&groups)
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot find latest record err: %v", err))
		return groups[0], err
	}
	return groups[0], nil
}
