package sfgroup

import (
	"fmt"
)

type IGroupRepository interface {
	Find()
	Insert()
	Removal()
	Update()
}

type GroupRepository struct {

}

func NewGroupRepository() *GroupRepository{

}