package server

import (
	"Forum-back/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartServer() {

	serverAddr := os.Getenv("BIND_ADDRESS") + ":" + os.Getenv("LISTEN_PORT")
	sslAvailable := true
	server := &http.Server{
		Addr: serverAddr,
	}

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
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Println("[HTTP] Received shutdown signal, shutting down HTTP server...")
		server.Shutdown(nil)
	}()

	if sslAvailable {
		log.Printf("[HTTPS] Starting HTTPS Server on https://%s\n", serverAddr)
		if err := server.ListenAndServeTLS(os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE")); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[HTTPS] Error occurs on HTTPS server : %v", err)
		}
		log.Println("[HTTPS] HTTPS server stop gracefully .")
	} else {
		log.Printf("[HTTP] Starting HTTP Server on http://%s\n", serverAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[HTTP] Error occurs on HTTP server : %v", err)
		}
		log.Println("[HTTP] HTTP server stop gracefully .")
	}

}
