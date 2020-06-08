package main

import (
	"clean-arch/models"
	"clean-arch/delivery/controllers"
	"clean-arch/delivery/protocols"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

const (
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 90 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	router := mux.NewRouter()

	ct := controllers.NewSignupController()
	router.HandleFunc("/signup", Adapt(ct)).Methods(http.MethodPost)
	server := &http.Server{
		Addr:              "127.0.0.1:5050",
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Printf("Server running on port %s\n", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func Adapt(c protocols.Controller) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := c.Handle(models.HttpRequest{
			Body: r.Body,
		})

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		w.Write(jsonResponse)
	}
}
