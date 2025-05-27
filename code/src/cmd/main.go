package main

import (
	"Forum-back/internal/config"
	"os"
)

func main() {
	config.LoadEnv()
	println(os.Getenv("GOOGLE_CLIENT_SECRET"))
}
