package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	log.Print("starting server...")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-nr-test-app"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", handler))

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
