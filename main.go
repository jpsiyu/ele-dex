package main

import (
	"ele-dex/api"
	"ele-dex/ws"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = 19100

func LogUri(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// log setting
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	hub := ws.NewHub()
	go hub.Run()

	r := mux.NewRouter()
	r.Use(LogUri)

	r.HandleFunc("/greet", api.Greet)
	r.HandleFunc("/echo", ws.Echo)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.NewClient(hub, w, r)
	})

	// cross domain
	optionHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	optionMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	optionOrigins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Server listening on port:", port)
	http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CORS(optionHeaders, optionMethods, optionOrigins)(r),
	)
}
