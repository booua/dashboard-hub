package firebasemodule

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
)

func GetOpeningTime() string{
	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://dashboard-hub.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	opening := ""

	if err := client.NewRef("blinds/openingTime").Get(ctx, &opening); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", opening)
	return opening
}

func GetClosingTime() string{

	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://dashboard-hub.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	closing := ""

	if err := client.NewRef("blinds/closingTime").Get(ctx, &closing); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", closing)
	return closing

}

func SetClosingTime(closingTime string) {
	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://dashboard-hub.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	closing := closingTime

	if err := client.NewRef("blinds/closingTime").Set(ctx, &closing); err != nil {
		log.Fatal(err)
	}
}

func SetOpeningTime(openingTime string) {
	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://dashboard-hub.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	opening := openingTime

	if err := client.NewRef("blinds/openingTime").Set(ctx, &opening); err != nil {
		log.Fatal(err)
	}
}

func SetupFirebase(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://dashboard-hub.firebaseio.com",
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	closing := ""
	opening := ""

	if err := client.NewRef("blinds/closingTime").Get(ctx, &closing); err != nil {
		log.Fatal(err)
	}

	if err := client.NewRef("blinds/openingTime").Get(ctx, &opening); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", opening)
	log.Printf("%s\n", closing)

}
