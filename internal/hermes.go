package hermes

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jbell-circle/hermes/internal/controllers"
	"github.com/jbell-circle/hermes/internal/views"
)

func Run(addr string) error {
	// controllers
	addrCtrl := controllers.NewAddressController()

	// views
	addrView := views.NewAddressView(addrCtrl)

	log.Printf("[INFO] server listening on %s\n", addr)
	router := NewRouter(addrView)
	return http.ListenAndServe(addr, router)
}

func NewRouter(addrView *views.AddressView) *httprouter.Router {
	router := httprouter.New()

	addrView.Register(router)

	return router
}
