package middleware

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"rate-limiter/internal/models"
)

var (
	clients = make(map[string]*models.Client)
	mu      sync.Mutex
)

const (
	Limit  = 5
	Window = time.Minute
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		mu.Lock()

		client, exists := clients[ip]

		if !exists {
			clients[ip] = &models.Client{
				Count:       1,
				WindowStart: time.Now(),
			}

			mu.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if time.Since(client.WindowStart) > Window {
			client.Count = 1
			client.WindowStart = time.Now()

			mu.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if client.Count >= Limit {
			mu.Unlock()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)

			json.NewEncoder(w).Encode(map[string]string{
				"error": "Rate limit exceeded",
			})

			return
		}

		client.Count++

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}