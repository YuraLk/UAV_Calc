package controller

import (
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	EnvTemp  float32 `json:"env_temp"`
	EnvPress float32 `json:"env_press"`
}

// Расчет характеристик
func Calculate(c *gin.Context) {

}
