package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// Getwd returns the current working directory path
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Working Directory:", path)
	// Find the path name for the active executable
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Extract the directory from the executable path
	exPath := filepath.Dir(ex)
	fmt.Println("Executable Directory:", exPath)

	router := gin.Default()
	router.GET("/caddi-cmd/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/caddi-cmd/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy",
		})
	})
	router.GET("/caddi-cmd/readyz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthy",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
	//router.RunTLS(":8443", "/caddi/cmd/server.pem", "/caddi/cmd/server-key.pem")
}
