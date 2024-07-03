package controllers

//go:generate mockgen -destination=../mocks/address_controller_mocks.go -package=mocks . AddressController

import (
	"github.com/jbell-circle/hermes/internal/models"
)

type AddressController interface {
	CreateAddress(id string, addr models.Address) (models.Address, error)
	GetAddress(id string) (models.Address, error)
}

func NewAddressController(addrDAO models.AddressDAO) AddressController {
	return &addressControllerImpl{
		addrDAO: addrDAO,
	}
}

type addressControllerImpl struct {
	addrDAO models.AddressDAO
}

func (ctrl *addressControllerImpl) CreateAddress(id string, addr models.Address) (models.Address, error) {
	return ctrl.addrDAO.Create(id, addr)
}

func (ctrl *addressControllerImpl) GetAddress(id string) (models.Address, error) {
	panic("implement me")
}
