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

	OverallApi := Api.Group("overall")
	{
		OverallApi.GET("", handler.GetOverall)
		OverallApi.GET("list")
	}

	AreaApi := Api.Group("area")
	{
		AreaApi.GET("provinceName/list")

		AreaApi.GET("province")
		AreaApi.GET("province/list")

		AreaApi.GET("city")
		AreaApi.GET("city/list")
	}

	AbroadApi := Api.Group("abroad")
	{
		AbroadApi.GET("")
		AbroadApi.GET("list")
	}

	TimelineApi := Api.Group("timeline")
	{
		TimelineApi.GET("")
		TimelineApi.GET("list")
	}
}
