package cron

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/booua/dashboard-hub/backend/services/mqtt"
	"gopkg.in/robfig/cron.v2"
)

type TimeSetup struct {
	CronExpression string
}

func SetupTimeForOpening(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var post TimeSetup
	err := decoder.Decode(&post)

	if err != nil {
		panic(err)
	}
	SetupCronJobForOpening(post.CronExpression)
}

func SetupTimeForClosing(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var post TimeSetup
	err := decoder.Decode(&post)

	if err != nil {
		panic(err)
	}
	SetupCronJobForClosing(post.CronExpression)
}

func SetupCronJobForOpening(cronExpression string) {
	c := cron.New()
	c.AddFunc(cronExpression, func() {
		mqtt.PerformBlindsAction("OPEN")
		fmt.Println("Opening the blinds at %s", cronExpression)
	})
	c.Start()
}

func SetupCronJobForClosing(cronExpression string) {
	c := cron.New()
	c.AddFunc(cronExpression, func() {
		mqtt.PerformBlindsAction("CLOSE")
		fmt.Println("Closing blinds at %s", cronExpression)
	})
	c.Start()
}
