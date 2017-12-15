package dao

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/astaxie/beego/orm"
	"marathon/redispool"
)

/**
*Author:hanyajun
*Time:17-12-15 上午11:03
*Discription:
**/
type Address struct {
	NodeId    int64  `json:"node_id" orm:"auto"`
	ParentId  int64  `json:"parent_id" orm:"column(parent_id)"`
	ChildId   int64  `json:"child_id" orm:"column(child_id)"`
	NodeValue string `json:"node_value" orm:"column(node_value)"`
}
type Province struct {
	City map[int64]City
}
type City struct {
	Area map[int64]string
}
type IAddressDao interface {
	SelectAll() ([]Address, error)
}
type AddressDao struct {
	redisP *redispool.RedisPool
	log    kitlog.Logger
}

func NewAddressDao(log kitlog.Logger, redisP *redispool.RedisPool) (*AddressDao, error) {
	addressDao := &AddressDao{
		log:    log,
		redisP: redisP,
	}
	orm.RegisterModel(new(Address))
	return addressDao, nil
}
func (addressDao *AddressDao) SelectAll() ([]Address, error) {
	o := orm.NewOrm()
	address := make([]Address, 0, 100)
	_, err := o.QueryTable("address").All(&address)
	//provices:=make([]*Province,0,100)
	//
	//for _,ad:=range *address{
	//	if ad.ParentId==""{
	//
	//		//if _, ok := province[ad.NodeId]; !ok{
	//	    	//
	//	    	//
	//		//}
	//		//province[ad.NodeId]=ad.NodeValue
	//	}
	//}
	return address, err
}
