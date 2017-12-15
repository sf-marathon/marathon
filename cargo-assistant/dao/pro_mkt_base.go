package dao

import (
	"github.com/astaxie/beego/orm"
	kitlog "github.com/go-kit/kit/log"
	"fmt"
)

const TABLE_NAME_PRO_MKT_BASE = "pro_market_base"

type ProMarketBase struct {
	MarketId         string  `json:"mkt_id" orm:"column(mkt_id);pk"`
	MarketNameShow   string  `json:"mkt_name_show" orm:"column(mkt_name_show)"`
	DailyMinPackages int     `json:"daily_min_packages" orm:"column(daily_min_packages)"`
	WeightMin        float64 `json:"weight_min" orm:"column(weight_min)"`
	WeightMax        float64 `json:"weight_max" orm:"column(weight_max)"`
	BasePrice        float64 `json:"base_price" orm:"column(base_price)"`
	BaseWeight       float64 `json:"base_weight" orm:"column(base_weight)"`
	GroupLimit       int     `json:"group_limit" orm:"column(group_limit)"`
	GroupDuration    int     `json:"group_duration" orm:"column(group_duration)"`
	UserRequire      string  `json:"user_require" orm:"column(user_require)"`
}
type IProMarketBaseDao interface {
	Select(id string) (*ProMarketBase, error)
	Insert(order *ProMarketBase) error
}
type ProMarketBaseDao struct {
	log kitlog.Logger
}

func NewProMarketBaseDao(log kitlog.Logger) (*ProMarketBaseDao, error) {
	orderDao := &ProMarketBaseDao{
		log: log,
	}
	return orderDao, nil
}

func (m *ProMarketBaseDao) Insert(r *ProMarketBase) error {
	//o := orm.NewOrm()
	//_, err := o.Insert(r)
	//if err != nil {
	//	m.log.Log("Cannot insert record %v, err: %v", r, err)
	//}
	return nil
}

func (m *ProMarketBaseDao) Select(id string) (*ProMarketBase, error) {
	proMktBase := ProMarketBase{MarketId: id}
	o := orm.NewOrm()
	err := o.QueryTable(TABLE_NAME_PRO_MKT_BASE).One(&proMktBase)
	if err != nil {
		m.log.Log(fmt.Sprintf("Cannot find record id %d, err: %v", id, err))
	}
	return &proMktBase, nil
}
