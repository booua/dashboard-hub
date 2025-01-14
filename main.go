package main

import (
	"log"
	"net/http"

	"github.com/booua/dashboard-hub/backend/services/cron"
	firebasemodule "github.com/booua/dashboard-hub/backend/services/firebase"
	"github.com/booua/dashboard-hub/backend/services/rest"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/dashboard-hub/blinds/", rest.BlindsActionsRoutes())
		r.Mount("/dashboard-hub/status/", rest.StatusRoutes())
		r.Mount("/dashboard-hub/time/", rest.SetupTimeRoutes())
	})
	return router
}

func main() {
	router := Routes()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}
	cron.SetupCronJobForClosing(firebasemodule.GetClosingTime())
	cron.SetupCronJobForOpening(firebasemodule.GetOpeningTime())
	log.Fatal(http.ListenAndServe(":8080", router))

}
