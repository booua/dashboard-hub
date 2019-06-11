package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var statusMsg = ""

func CloseBlinds(w http.ResponseWriter, r *http.Request) {
	PerformBlindsAction("CLOSE")
}

func OpenBlinds(w http.ResponseWriter, r *http.Request) {
	PerformBlindsAction("OPEN")
}

func PerformBlindsAction(action string) {
	const TOPIC = "node/dashboard"

	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.28:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("error on connecting")
	}
	if token := client.Publish(TOPIC, 0, false, action); token.Wait() && token.Error() != nil {
		fmt.Printf("error on publishing")
	}
}

func GetBlindsStatus(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)
	const TOPIC = "blinds/status"
	fetchMqttMessage(ctx, ch, w, r, TOPIC)
	select {
	case result := <-ch:
		fmt.Println(result)
		cancel()
		return
	}
	cancel()

	<-ch
}

func fetchMqttMessage(ctx context.Context, ch chan<- string, w http.ResponseWriter, r *http.Request, topic string) {

	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.28:1883")

	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			mqttMessageHandler(w, r, topic, msg, ch)
		}); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func GetMqttStatus(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)
	const TOPIC = "blinds/health"
	fetchMqttMessage(ctx, ch, w, r, TOPIC)
	select {
	case result := <-ch:
		fmt.Println(result)
		cancel()
		return
	}
	cancel()
	<-ch
}

type Status struct {
	Name   string    `json:"name"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

func mqttMessageHandler(w http.ResponseWriter, r *http.Request, statusName string, msg mqtt.Message, ch chan<- string) {
	statusMsg = string(msg.Payload())

	var status *Status = &Status{
		Name:   statusName,
		Status: statusMsg,
	}

	jsonStatus, _ := json.Marshal(status)

	w.Write(jsonStatus)
	ch <- statusMsg

}
