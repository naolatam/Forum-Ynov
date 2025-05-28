package main

import (
	"Forum-back/internal/config"
	hostedservices "Forum-back/pkg/hostedServices"
)

func main() {
	config.LoadEnv()
	go hostedservices.StartAllHostedServices()
}
