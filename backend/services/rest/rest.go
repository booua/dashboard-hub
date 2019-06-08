package rest

import (
	"fmt"
	"net/http"

	"github.com/booua/dashboard-hub/backend/services/mqtt"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/open", mqtt.OpenBlinds)
	router.Get("/close", mqtt.CloseBlinds)
	router.Get("/check", hello)
	return router
}

func hello(w http.ResponseWriter, r *http.Request) {

	// const TOPIC = "node/dashboard"

	// opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.28:1883")

	// client := mqtt.NewClient(opts)
	// if token := client.Connect(); token.Wait() && token.Error() != nil {
	// 	fmt.Printf("error on connecting")
	// }
	// if token := client.Publish(TOPIC, 0, false, "CLOSE"); token.Wait() && token.Error() != nil {
	// 	fmt.Printf("error on publishing")
	// }
	fmt.Fprintf(w, "hello")
}
