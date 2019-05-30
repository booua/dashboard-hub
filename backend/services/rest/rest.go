package rest

import (
	"github.com/booua/dashboard-hub/backend/services/mqtt"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/open", mqtt.OpenBlinds)
	router.Get("/close", mqtt.CloseBlinds)
	return router
}
