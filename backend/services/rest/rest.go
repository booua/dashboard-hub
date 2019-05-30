package rest

import (
	"/backend/services/mqtt/mqtt"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/open", mqtt.OpenBlinds)
	router.Get("/close", mqtt.CloseBlinds)
	return router
}
