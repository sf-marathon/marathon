package service

import (
	"marathon/cargo-assistant/dao"
)

/**
*Author:hanyajun
*Time:17-12-15 下午2:31
*Discription:
**/

type IAddressService interface {
	GetAllAddress() ([]dao.Address, error)
}

type AddressService struct {
	addressDao      dao.IAddressDao
}

func NewAddressService(address dao.IAddressDao) *AddressService {
	return &AddressService{
		addressDao:address,
	}
}
func (s *AddressService) GetAllAddress() ([]dao.Address, error) {

	return s.addressDao.SelectAll()
}

