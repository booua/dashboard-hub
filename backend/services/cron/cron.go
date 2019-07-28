package cron

import (
	"encoding/json"
	"fmt"
	"net/http"

	firebasemodule "github.com/booua/dashboard-hub/backend/services/firebase"
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
	firebasemodule.SetOpeningTime(post.CronExpression)
	SetupCronJobForOpening(post.CronExpression)
}

func SetupTimeForClosing(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var post TimeSetup
	err := decoder.Decode(&post)

	if err != nil {
		panic(err)
	}
	firebasemodule.SetClosingTime(post.CronExpression)
	SetupCronJobForClosing(post.CronExpression)
}

func SetupCronJobForOpening(cronExpression string) {
	c := cron.New()
	c.AddFunc(cronExpression, func() {
		mqtt.PerformBlindsAction("OPE")
		fmt.Println("Opening the blinds at %s", cronExpression)
	})
	c.Start()
}

func SetupCronJobForClosing(cronExpression string) {
	c := cron.New()
	c.AddFunc(cronExpression, func() {
		mqtt.PerformBlindsAction("CLOS")
		fmt.Println("Closing blinds at %s", cronExpression)
	})
	c.Start()
}
