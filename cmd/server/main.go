package main

import (
	"log"
	"net/http"

	"rate-limiter/internal/handlers"
	"rate-limiter/internal/middleware"
)

func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/",handlers.Home)
	mux.HandleFunc("/ping",handlers.Ping)

	wrappedMux := middleware.RateLimiter(mux)

	log.Println("server starting on : 8080")
	log.Fatal(http.ListenAndServe(":8080",wrappedMux))
}