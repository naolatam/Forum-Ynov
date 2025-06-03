package main

import (
	"Forum-back/internal/config"
	"Forum-back/internal/server"
	hostedservices "Forum-back/pkg/hostedServices"
	"Forum-back/pkg/repositories"
	"github.com/google/uuid"
	"log"
)

func main() {
	config.LoadEnv()
	go hostedservices.StartAllHostedServices()

	// Initialisation du repository
	userRepo := &repositories.UserRepository{}
	if !userRepo.Init() {
		log.Fatal("Échec de l'initialisation du UserRepository")
	}
	defer userRepo.Close()

	// Utilisation des valeurs existantes dans la base de données
	id := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000") // UUID de l'utilisateur inséré
	pseudo := "testPseudo"
	email := "test@example.com"

	// Test de la fonction FindByIdOrUsernameOrEmail
	user, err := userRepo.FindByIdOrUsernameOrEmail(id, pseudo, email)
	if err != nil {
		log.Printf("Erreur : %v\n", err)
	} else {
		log.Printf("Utilisateur trouvé : %+v\n", user)
	}
}
