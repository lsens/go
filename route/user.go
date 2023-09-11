package route

import (
	"github.com/gin-gonic/gin"
	"lss/contreller"
	"net/http"
)

func InitUser(engine *gin.Engine) {

	group := engine.Group("user")

	group.GET("info", func(c *gin.Context) {
		c.JSON(http.StatusOK, contreller.User.Info(c))
	})

}
