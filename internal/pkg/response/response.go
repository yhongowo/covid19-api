package response

import "github.com/gin-gonic/gin"

func ErrorByCode(c *gin.Context, code int, msg string) {
	c.JSON(code, msg)
	c.Abort()
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, data)
	c.Abort()
}
