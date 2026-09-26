package handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":"welcome to Rate Limiter API",
	})
}

func Ping(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":"pong",
	})
}