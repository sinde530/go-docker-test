package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pingTest)
	addr := "/8080"

	fmt.Printf("Server is listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func pingTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
