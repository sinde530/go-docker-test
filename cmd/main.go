package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinde530/go-docekr-test/internal/db"
)

func main() {
	client, err := db.ConnectToDB()
	if err != nil {
		fmt.Println("Failed to connected MongoDB", err.Error())
		return
	}

	defer client.Disconnect(context.Background())

	r := gin.Default()
	r.GET("/ping", pingTest)
	addr := "/8080"

	fmt.Printf("Server is listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func pingTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Rasberry Pi 4 Server Open."})
}
