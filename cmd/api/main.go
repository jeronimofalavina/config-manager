package main

import (
	"log"
	"os"

	route "HFtest-platform-jeronimofalavina-tst/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatal("SERVER_PORT environment variable is not defined.")
	}

	log.Printf("Server is running on port %s", port)
	route.SetupRoutes(r)

	r.Run(":" + port)
}
