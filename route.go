package main

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "lss/docs"
	"lss/middleware"
	"lss/model"
	"lss/route"
	"net/http"
)

func setupRouter() *gin.Engine {
	engine := gin.Default()

	//请求日志
	engine.Use(middleware.RequestLogger())

	//添加swagger支持
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RError(404, "无效的路由"))
	})

	engine.GET("test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "is test")
	})

	route.InitUser(engine)

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "系统异常，请联系客服",
		"code":    -1001,
	})
}
