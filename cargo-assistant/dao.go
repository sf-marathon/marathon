package cargo_assistant

import (
	"github.com/astaxie/beego/orm"
	kitlog "github.com/go-kit/kit/log"
)

/**
*Author:hanyajun
*Time:12/14/17 3:56 PM
*Discription:
**/

const TABLE_NAME_ORDER = "order"

type Order struct {
	Id              string       `json:"id" orm:"auto"`
	SenderName      string    `json:"sender_name" orm:"column(sender_name)"`
	SenderPhone     string    `json:"sender_phone" orm:"column(sender_phone)"`
	SenderAddress   string    `json:"sender_address" orm:"column(sender_address)"`
	ReceiverName    string    `json:"receiver_name" orm:"column(receiver_name)"`
	ReceiverPhone   string    `json:"receiver_phone" orm:"column(receiver_phone)"`
	ReceiverAddress string    `json:"receiver_address" orm:"column(receiver_address)"`
	CreateTime      string `json:"create_time" orm:"column(create_time)"`
	Weight          int       `json:"weight" orm:"column(weight)"`
	Type            string    `json:"type" orm:"column(type)"`
	Fee             int       `json:"fee" orm:"column(fee)"`
	ExpectTime      string `json:"expect_time" orm:"column(expect_time)"`
	Labels          string    `json:"labels" orm:"column(labels)"`
	Status          int       `json:"status" orm:"column(status)"`
}
type IOrderDao interface {
	Select(id string) (*Order, error)
	Insert(order *Order) error
	Cancel(id string) error
}
type OrderDao struct {
	log kitlog.Logger
}

func NewOrderDao(log kitlog.Logger) (*OrderDao, error) {
	orderDao := &OrderDao{
		log: log,
	}
	return orderDao, nil
}

func (m *OrderDao) Insert(r *Order) error {
	//o := orm.NewOrm()
	//_, err := o.Insert(r)
	//if err != nil {
	//	m.log.Log("Cannot insert record %v, err: %v", r, err)
	//}
	return nil
}

func (m *OrderDao) Select(id string) (*Order, error) {
	//intId, err := strconv.Atoi(id)
	//if err != nil {
	//	return nil, err
	//}
	order := Order{Id:id}
	//o := orm.NewOrm()
	//err := o.QueryTable(TABLE_NAME_ORDER).One(&order)
	//if err != nil {
	//	m.log.Log("Cannot insert record id %d, err: %v", id, err)
	//}
	return &order, nil
}

func (m *OrderDao) Cancel(id string) error {
	//intId, err := strconv.Atoi(id)
	//if err != nil {
	//	return err
	//}
	order := Order{Id:id, Status:0}
	o := orm.NewOrm()
	count, err := o.Update(order)
	if err != nil {
		m.log.Log("Error when cancel order %d", id)
	}
	m.log.Log("Cancelled %d order for id %d",count,id)
	return nil
}
