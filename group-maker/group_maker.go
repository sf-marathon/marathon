package group_maker

import (
	ca "marathon/cargo-assistant"
	"marathon/cargo-assistant/dao"
)

type IGroupMaker interface {}

type GroupMaker struct {
	orderQ chan dao.order
	dao dao.IGroupDao
}

func NewGroupMaker(q chan *ca.Order, dao dao.IGroupDao) *GroupMaker {
	return &GroupMaker{
		orderQ: q,
		dao:dao,
	}
}

func (g *Group)Run(){
	for{
		select{
		case order <- g.orderQ:
			//svc send order here
			//store into db
			//combination group
		}
	}
}

func (g *Group)combination(order){
	//用g.dao取到现有团
	//进行join团or new group判断

	//if join

	//if new, store a new group with g.dao
}
