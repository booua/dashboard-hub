package rest

import (
	"github.com/booua/dashboard-hub/backend/services/cron"
	"github.com/booua/dashboard-hub/backend/services/mqtt"

	"github.com/go-chi/chi"
)

func BlindsActionsRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/open", mqtt.OpenBlinds)
	router.Get("/close", mqtt.CloseBlinds)
	return router
}

func StatusRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/blinds-status", mqtt.GetBlindsStatus)
	router.Get("/mqtt-status", mqtt.GetMqttStatus)
	return router
}

func SetupTimeRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/closing", cron.SetupTimeForClosing)
	router.Post("/opening", cron.SetupTimeForOpening)
	return router
}
