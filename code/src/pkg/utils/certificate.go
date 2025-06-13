package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

// CheckIfCertExist checks if the certificate and key files exist.
func CheckIfCertExist(certFile, keyFile string) bool {

	// Vérifie si les fichiers existent, sinon les génère
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Println("[HTTPS] Unable to find certificate files.")
		return false
	}
	return true

}

// GenerateSelfSignedCert generates a self-signed certificate and saves it to the specified files.
func GenerateSelfSignedCert(certFile, keyFile string, length int) error {
	priv, err := rsa.GenerateKey(rand.Reader, length)
	if err != nil {
		return err
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return err
	}

	// Create a self-signed certificate template (CA)
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Forum-Ynov"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Add localhost as a DNS name and IP address
	template.DNSNames = []string{"localhost"}
	template.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	// Write the certificate
	certOut, err := os.Create(certFile)
	if err != nil {
		return err
	}
	defer certOut.Close()
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// Write the private key
	keyOut, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer keyOut.Close()
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return nil
}
