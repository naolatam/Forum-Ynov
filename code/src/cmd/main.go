package main

import (
	"Forum-back/internal/config"
	"Forum-back/internal/server"
	hostedservices "Forum-back/pkg/hostedServices"
)

func main() {
	config.LoadEnv()
	go hostedservices.StartAllHostedServices()

	server.StartServer()
}
