package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinde530/go-docekr-test/cmd/cpu"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pingTest)
	// r.POST("/signup", HandleSignup)
	r.GET("/cpu", cpu.GetCPUTemperature)
	addr := "/8080"

	fmt.Printf("Server is listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func pingTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Rasberry Pi 4 Server Open."})
}
