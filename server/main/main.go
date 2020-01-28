/*
 * Generated file main.go.  Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */

package main

import (
  "context"
  "flag"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/jaekwon.park/petstore-go/endpoints/pet"
  "github.com/jaekwon.park/petstore-go/endpoints/utils"
  "github.com/jaekwon.park/petstore-go/server/routes"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"
)

func main() {
  var wait time.Duration
  flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
  flag.Parse()

  r := mux.NewRouter()

  // SETUP THE ROUTE MAPPINGS
  for k, v := range routes.HandlerMap() {
    r.HandleFunc(k.Path, v).Methods(k.Method)
  }

  srv := &http.Server{
    Addr:         "0.0.0.0:8082",
    // Good practice to set timeouts to avoid Slowloris attacks.
    WriteTimeout: time.Second * 15,
    ReadTimeout:  time.Second * 15,
    IdleTimeout:  time.Second * 60,
    Handler: r, // Pass our instance of gorilla/mux in.
  }

  // Run our server in a goroutine so that it doesn't block.
  go func() {
    if err := srv.ListenAndServe(); err != nil {
      log.Println(err)
    }
  }()

  c := make(chan os.Signal, 1)
  // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
  // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
  signal.Notify(c, os.Interrupt)

  // Block until we receive our signal.
  <-c

  // Create a deadline to wait for.
  ctx, cancel := context.WithTimeout(context.Background(), wait)
  defer cancel()
  // Doesn't block if no connections, but will otherwise wait
  // until the timeout deadline.
  srv.Shutdown(ctx)
  // Optionally, you could run srv.Shutdown in a goroutine and block on
  // <-ctx.Done() if your application should wait for other services
  // to finalize based on context cancellation.
  log.Println("shutting down")
  os.Exit(0)
}

func init() {
  // Set client options
  clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

  // Connect to MongoDB
  client, err := mongo.Connect(context.TODO(), clientOptions)

  if err != nil {
    log.Fatal(err)
  }

  // Check the connection
  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Connected to MongoDB!")
  mc := client.Database("test").Collection("pet")

   // SET THE IMPLEMENTATIONS FOR EACH SERVICE HERE.  THERE SHOULD BE ONE ENTRY FOR EACH
   // NAMESPACE SERVED BY THIS SERVER.
  routes.SetServiceImplementation(routes.PET_API_IMPL_WRAPPER, pet.Service{Mc:mc})
  routes.SetServiceImplementation(routes.UTILS_API_IMPL_WRAPPER, utils.Service{})

   //  THIS METHOD IS NEEDED TO ENSURE THAT AS APIS EVOLVE THE REST SERVER FAILS TO STARTUP
   // ALLOWING US TO CATCH ERRORS EARLY
  // validate that all endpoints are valid
  routes.ValidateReadinessToRun()
}