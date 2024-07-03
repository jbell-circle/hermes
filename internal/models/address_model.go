package models

import (
	"errors"
	"sync"
)

//go:generate mockgen -destination=../mocks/address_model_mocks.go -package=mocks . AddressDAO

type Address struct {
	Name     string
	Address1 string
	Address2 string
	City     string
	State    string
	ZipCode  string
}

type AddressDAO interface {
	Create(id string, addr Address) (Address, error)
	Get(id string) (Address, error)
}

func NewAddressDAO() AddressDAO {
	return new(addressDAOMemory)
}

type addressDAOMemory struct {
	store sync.Map
}

func (dao *addressDAOMemory) Create(id string, addr Address) (Address, error) {
	val, _ := dao.store.LoadOrStore(id, addr)
	retAddr, ok := val.(Address)
	if !ok {
		return Address{}, errors.New("unexpected type returned from addressDAOMemory store")
	}
	return retAddr, nil
}

func (dao *addressDAOMemory) Get(id string) (Address, error) {
	val, _ := dao.store.Load(id)
	retAddr, ok := val.(Address)
	if !ok {
		return Address{}, errors.New("address not found in addressDAOMemory store")
	}
	return retAddr, nil
}
