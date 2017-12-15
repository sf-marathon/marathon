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
	//TODO: find last group of this base
	count,err := o.QueryTable(TABLE_NAME_GROUP).Count()
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot get total group info err: %v", err))
		return err
	}
	now := time.Now()
	count = count%100
	suffix := fmt.Sprintf("%s%02d",now.Format("0102"),count)
	fmt.Println(proMktBase.MarketNameShow + suffix)
	group := Group{
		MarketId:proMktBase.MarketId,
		GroupName:proMktBase.MarketNameShow + suffix,//TODO: append time info
		CreateTime:time.Now(),
		DueTime:now.Add(60*time.Second *time.Duration(proMktBase.GroupDuration)),
		}
	_, err = o.Insert(&group)
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
