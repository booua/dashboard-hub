package firebasemodule

import (
  "firebase.google.com/go"
  "net/http"
  "context"
)

func TestFirebase(w http.ResponseWriter, r *http.Request) {

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
  
  acc := OperationTime{
    closingTime: "TEST",
  }

    if err := client.NewRef("blinds/operationTimes").Set(ctx, acc); err != nil {
      log.Fatal(err)
  }
}
