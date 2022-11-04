package router

import (
	"covid19-api/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"

	"net/http"
)

func SetUp(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
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
