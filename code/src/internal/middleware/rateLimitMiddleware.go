package middleware

import (
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type client struct {
	requests             int
	requestsOnStaticPage int
	lastSeen             time.Time
}

var (
	clients     = make(map[string]*client)
	clientsLock sync.Mutex
)

const (
	maxStaticRequests  = 80
	maxGeneralRequests = 30
)

// RateLimitMiddleware limits the number of requests a client can make to the server.
func RateLimitMiddleware(next http.Handler) http.Handler {
	if next == nil {
		next = http.DefaultServeMux
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := extractIP(r.RemoteAddr)
		isStatic := isStaticRequest(r)
		now := time.Now()

		if exceeded := handleClientRateLimit(ip, isStatic, now); exceeded {
			log.Printf("[RateLimit] %s exceeded limit on %s\n", ip, r.URL.Path)
			w.Header().Set("Retry-After", "60")
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// extractIP extracts the IP address from the remote address string.
func extractIP(remoteAddr string) string {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return host
}

// isStaticRequest checks if the request is for a static resource.
func isStaticRequest(r *http.Request) bool {
	return strings.HasPrefix(r.URL.Path, "/static")
}

// handleClientRateLimit checks and updates the client's request count.
func handleClientRateLimit(ip string, isStatic bool, now time.Time) (rateLimitExceeded bool) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	c, exists := clients[ip]
	if !exists || now.Sub(c.lastSeen) > time.Minute {
		clients[ip] = &client{lastSeen: now}
		if isStatic {
			clients[ip].requestsOnStaticPage = 1
		} else {
			clients[ip].requests = 1
		}
		return false
	}

	if c.requestsOnStaticPage >= maxStaticRequests || c.requests >= maxGeneralRequests {
		return true
	}

	if isStatic {
		c.requestsOnStaticPage++
	} else {
		c.requests++
	}
	c.lastSeen = now
	return false
}
