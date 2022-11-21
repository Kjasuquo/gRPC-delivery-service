package main

import (
	// "log"
	myService "gitlab.com/dh-backend/delivery-service/internal/services"
	// "github.com/joho/godotenv"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("Error loading .env with godotenv: %s", err)
	// }
	myService.Start()
}
