package handler

import (
	"covid19-api/internal/pkg/response"
	"covid19-api/internal/service"
	"github.com/gin-gonic/gin"
)

func GetOverall(c *gin.Context) {
	overallService := service.NewOverallService(c)
	result, err := overallService.GetLatest()
	if err != nil {
		response.ErrorByCode(c, 200, err.Error())
	}
	response.Success(c, result)
}
