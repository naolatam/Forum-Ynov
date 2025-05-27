package middleware

import (
	"net/http"
	"sync"
	"time"
)

type client struct {
	requests int
	lastSeen time.Time
}

var (
	clients     = make(map[string]*client)
	clientsLock sync.Mutex
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		clientsLock.Lock()
		c, exists := clients[ip]
		if !exists || time.Since(c.lastSeen) > time.Minute {
			clients[ip] = &client{requests: 1, lastSeen: time.Now()}
			clientsLock.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if c.requests >= 10 {
			clientsLock.Unlock()
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		c.requests++
		c.lastSeen = time.Now()
		clientsLock.Unlock()
		next.ServeHTTP(w, r)
	})
}
