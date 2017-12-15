package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	kitlog "github.com/go-kit/kit/log"
)

const (
	TABLE_NAME_JOIN = "join"

	COLUMN_JOIN_ID       = "join_id"
	COLUMN_JOIN_GROUP_ID = "group_id"
	COLUMN_PHONE         = "phone"
)

type Join struct {
	JoinId            int  `json:"join_id" orm:"column(join_id);pk;auto"`
	GroupId           int     `json:"group_id" orm:"column(group_id)"`
	Phone             string  `json:"phone" orm:"column(phone)"`
	Address           string  `json:"address" orm:"column(address)"`
	ExpectDailyAmount int     `json:"expect_daily_amount" orm:"column(expect_daily_amount)"`
	TotalAmount       int     `json:"total_amount" orm:"column(total_amount)"`
	AverageWeight     float64 `json:"average_weight" orm:"column(average_weight)"`
}

type IJoinDao interface {
	Select(groupId string) ([]*Join, error)
	Insert(joinReq *Join) error
}

type JoinDao struct {
	log kitlog.Logger
}

func NewJoinDao(log kitlog.Logger) (*JoinDao, error) {
	joinDao := &JoinDao{
		log: log,
	}
	return joinDao, nil
}

func (m *JoinDao) Insert(join *Join) error {
	o := orm.NewOrm()
	var joined *Join
	//check if the user is already in this group
	err := o.QueryTable(TABLE_NAME_JOIN).One(joined, COLUMN_JOIN_GROUP_ID, fmt.Sprintf("%d", join.GroupId), COLUMN_PHONE, join.Phone)
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot comfirm user record err: %v", err))
		return err
	} else if join != nil {
		m.log.Log(fmt.Sprintf("User already joined"))
		return fmt.Errorf("user already joined group %d", join.GroupId)
	}
	_, err = o.Insert(join)
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot insert record err: %v", err))
		return err
	}
	return nil
}

func (m *JoinDao) Select(groupId string) ([]*Join, error) {
	var joins []*Join
	o := orm.NewOrm()
	_, err := o.QueryTable(TABLE_NAME_JOIN).Filter(COLUMN_JOIN_GROUP_ID, groupId).All(&joins)
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot find record group %s: %b", groupId, err))
		return joins, err
	}
	return joins, nil
}
