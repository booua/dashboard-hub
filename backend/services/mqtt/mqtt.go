package mqtt

import (
	"fmt"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CloseBlinds(w http.ResponseWriter, r *http.Request) {

	const TOPIC = "node/dashboard"

	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.28:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("error on connecting")
	}
	if token := client.Publish(TOPIC, 0, false, "CLOSE"); token.Wait() && token.Error() != nil {
		fmt.Printf("error on publishing")
	}
}

func OpenBlinds(w http.ResponseWriter, r *http.Request) {

	const TOPIC = "node/dashboard"

	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.28:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("error on connecting")
	}
	if token := client.Publish(TOPIC, 0, false, "OPEN"); token.Wait() && token.Error() != nil {
		fmt.Printf("error on publishing")
	}
}
