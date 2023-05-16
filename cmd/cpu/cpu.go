package cpu

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCPUTemperature(c *gin.Context) {
	temperature, err := readCPUTemperature()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"temperature": temperature})
}

func readCPUTemperature() (float32, error) {
	data, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	temperature := strings.TrimSpace(string(data))
	temperatureInt, err := strconv.Atoi(temperature)
	if err != nil {
		return 0, err
	}

	temperatureCelsius := float32(temperatureInt) / 1000
	return temperatureCelsius, nil
}
