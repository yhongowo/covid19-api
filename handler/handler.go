package handler

import (
	"covid19-api/service"
	"github.com/gin-gonic/gin"
)

func GetOverall(c *gin.Context) {
	overallService := service.NewOverallService(c)
	result, err := overallService.GetLatest()
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, result)
}
