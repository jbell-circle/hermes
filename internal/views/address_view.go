package views

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jbell-circle/hermes/internal/controllers"
	"github.com/jbell-circle/hermes/internal/models"

	"github.com/julienschmidt/httprouter"
)

const (
	pathPostAddresses = "/addresses/:id"
)

type AddressView struct {
	addrCtrl controllers.AddressController
}

func NewAddressView(addrCtrl controllers.AddressController) *AddressView {
	return &AddressView{
		addrCtrl: addrCtrl,
	}
}

func (view *AddressView) Register(router *httprouter.Router) {
	router.POST(pathPostAddresses, view.CreateAddress)
}

func (view *AddressView) CreateAddress(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	var addr models.Address
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&addr); err != nil {
		log.Printf("[ERROR] failed to decode request body: %s\n", err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	addr, err := view.addrCtrl.CreateAddress(id, addr)
	if err != nil {
		log.Printf("[ERROR] failed to create address (%s): %s\n", id, err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	if err := enc.Encode(addr); err != nil {
		log.Printf("[ERROR] failed to encode response body: %s\n", err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	log.Printf("[DEBUG] created address (%s)\n", id)
}
