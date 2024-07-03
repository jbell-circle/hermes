package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressDAOMemory_Create(t *testing.T) {
	addrDAO := NewAddressDAO()

	// new key, returns new entry
	id, addr := "circle", Address{
		Name:     "Circle",
		Address1: "99 High St",
		City:     "Boston",
		State:    "MA",
		ZipCode:  "02110",
	}
	retAddr, err := addrDAO.Create(id, addr)
	assert.NoError(t, err)
	assert.Equal(t, addr, retAddr)

	// duplicate key, returns original entry
	newAddr := Address{
		Name:     "Circle",
		Address1: "99 High St",
		City:     "Boston",
		State:    "MA",
		ZipCode:  "02110",
	}
	retAddr, err = addrDAO.Create(id, newAddr)
	assert.NoError(t, err)
	assert.Equal(t, addr, retAddr)
}
