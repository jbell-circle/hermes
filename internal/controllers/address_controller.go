package controllers

//go:generate mockgen -destination=../mocks/address_controller_mocks.go -package=mocks . AddressController

import (
	"log"

	"github.com/jbell-circle/hermes/internal/models"
)

type AddressController interface {
	CreateAddress(id string, addr models.Address) (models.Address, error)
}

func NewAddressController() AddressController {
	return &addressControllerImpl{}
}

type addressControllerImpl struct{}

func (ctrl *addressControllerImpl) CreateAddress(id string, addr models.Address) (models.Address, error) {
	// TODO: implement me
	log.Printf("[WARN] implement me")
	return addr, nil
}
