package cargo_assistant

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/astaxie/beego/orm"
)
/**
*Author:hanyajun
*Time:17-12-15 上午11:03
*Discription:
**/
type Address struct {
	NodeId  int64  `json:"node_id" orm:"auto"`
	ParentId string `json:"parent_id" orm:"column(parent_id)"`
	ChildId  string `json:"parent_id" orm:"column(parent_id)"`
	NodeValue string `json:"node_value" orm:"column(parent_id)"`
}
type Province struct {
	City map[int64] City
}
type City struct {
	Area map[string]string
}
type IAddressDao interface {
	SelectAll() ([]*Province, error)
}
type AddressDao struct {
	log kitlog.Logger
}

func NewAddressDao(log kitlog.Logger) (*AddressDao, error) {
	addressDao := &AddressDao{
		log: log,
	}
	orm.RegisterModel(new(Address))
	return addressDao, nil
}
//func SelectAll()([]*Province,error ) {
//	//o := orm.NewOrm()
//	//address:=new([]Address)
//	//_,err:=o.QueryTable("address").All(address)
//	//provices:=make([]*Province,0,100)
//	//province := make(map[int64]City)
//	//for _,ad:=range *address{
//	//	if ad.ParentId==""{
//	//
//	//		province:=append(province,)
//	//		//if _, ok := province[ad.NodeId]; !ok{
//	//	    	//
//	//	    	//
//	//		//}
//	//		//province[ad.NodeId]=ad.NodeValue
//	//	}
//	//}
//	return ,err
//}