package views

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	"github.com/jbell-circle/hermes/internal/mocks"
	"github.com/jbell-circle/hermes/internal/models"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	addrCtrl := mocks.NewMockAddressController(ctrl)

	addrView := NewAddressView(addrCtrl)

	router := httprouter.New()
	addrView.Register(router)

	address := models.Address{
		Name:     "Circle",
		Address1: "99 High St",
		City:     "Boston",
		State:    "MA",
		ZipCode:  "02110",
	}
	addrCtrl.EXPECT().CreateAddress(gomock.Eq("circle"), gomock.Eq(address))

	buf, _ := json.Marshal(address)
	req := httptest.NewRequest("POST", "/addresses/circle", bytes.NewReader(buf))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusCreated)
}

func TestBar(t *testing.T) {
	ctrl := gomock.NewController(t)
	addrCtrl := mocks.NewMockAddressController(ctrl)

	addrView := NewAddressView(addrCtrl)

	router := httprouter.New()
	addrView.Register(router)

	address := models.Address{
		Name:     "Circle",
		Address1: "99 High St",
		City:     "Boston",
		State:    "MA",
		ZipCode:  "02110",
	}
	addrCtrl.EXPECT().
		GetAddress(gomock.Eq("circle")).
		Return(address, nil)

	req := httptest.NewRequest("GET", "/addresses/circle", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}
