package handler

import (
	"covid19-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOverall(c *gin.Context) {
	overallService := service.NewOverallService(c)
	result, err := overallService.GetLatest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, result)
}