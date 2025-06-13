package server

import (
	"Forum-back/internal/middleware"
	"Forum-back/internal/routes"
	"Forum-back/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// StartServer initializes the server and starts listening for incoming requests.
func StartServer() {

	serverAddr := os.Getenv("BIND_ADDRESS") + ":" + os.Getenv("LISTEN_PORT")
	sslAvailable := true
	server := &http.Server{
		Addr:    serverAddr,
		Handler: middleware.RateLimitMiddleware(nil),
	}

	// Initialize routes
	routes.InitRoutes()

	// Check if the certificate and key files exist
	if !utils.CheckIfCertExist(os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE")) {
		log.Printf("[HTTP] No certificate found, generating self-signed certificate for HTTPS on %s\n", serverAddr)
		if err := utils.GenerateSelfSignedCert(os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"), 2048); err != nil {
			sslAvailable = false
			log.Fatalf("[HTTPS] Failed to generate self-signed certificate: %v", err)
		}
	} else {
		log.Printf("[HTTPS] Certificate found, starting HTTPS server on %s\n", serverAddr)
	}

	// Allow graceful shutdown
	go gracefulShutdown(server)

	// Start HTTPS or HTTP server based on certificate availability
	if sslAvailable {
		startHTTPSServer(server)
	} else {
		startHTTPServer(server)
	}

}

// startHTTPSServer starts the HTTPS server with the provided server configuration.
func startHTTPSServer(server *http.Server) {
	log.Printf("[HTTPS] Starting HTTPS Server on https://%s\n", server.Addr)
	if err := server.ListenAndServeTLS(os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE")); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[HTTPS] Error occurs on HTTPS server : %v", err)
	}
	log.Println("[HTTPS] HTTPS server stop gracefully .")
}

// startHTTPServer starts the HTTP server with the provided server configuration.
func startHTTPServer(server *http.Server) {
	log.Printf("[HTTP] Starting HTTP Server on http://%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[HTTP] Error occurs on HTTP server : %v", err)
	}
	log.Println("[HTTP] HTTP server stop gracefully .")
}

// gracefulShutdown listens for OS signals to gracefully shut down the server.
func gracefulShutdown(server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("[HTTP] Received shutdown signal, shutting down server...")
	if err := server.Shutdown(nil); err != nil {
		log.Fatalf("[HTTP] Error during server shutdown: %v", err)
	}
	log.Println("[HTTP] Server stopped gracefully.")
}
