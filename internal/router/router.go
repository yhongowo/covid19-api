package router

import (
	"covid19-api/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp(r *gin.Engine) {
	r.Use(gin.Recovery())

	Api := r.Group("api")
	{
		Api.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome covid19-api")
		})

	}

	Api.GET("overall", handler.GetOverall)
	Api.GET("overall/list", handler.ListOverall)

	AreaApi := Api.Group("area")
	{
		AreaApi.GET("abroad", handler.ListAbroad)

		AreaApi.GET("provinceName", handler.ListProvinceName)
		AreaApi.GET("provinceShortName", handler.ListProvinceShortName)
		AreaApi.GET("province", handler.ListProvince)
		AreaApi.GET("province/list")

		AreaApi.GET("city")
		AreaApi.GET("city/list")
	}

	Api.GET("timeline", handler.ListTimeline)
}
