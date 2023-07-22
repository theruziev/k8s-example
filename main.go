package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello from k8s example - master\n"
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println("Error getting hostname:", err)
			return
		}

		// Resolve the IP addresses associated with the hostname
		addresses, err := net.LookupIP(hostname)
		if err != nil {
			fmt.Println("Error looking up IP addresses:", err)
			return
		}
		msg += fmt.Sprintf("IPs: %s", addresses)
		w.Write([]byte(msg))
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Get("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
