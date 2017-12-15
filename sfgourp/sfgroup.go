package sfgroup

import (
	"fmt"
	ca "marathon/cargo-assistant"
)

type IGroup interface{

}

type Group struct {
	orderQ chan *ca.Order
	repository IGroupRepository
}

func NewGroup(q chan *ca.Order, repository IGroupRepository) *Group {
	return &Group{
		orderQ: q,
		repository:repository,
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

func (g *Group)
