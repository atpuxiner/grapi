// 划分版本号，便于后期扩展（v|V）

package router

import (
	"github.com/atpuxiner/grapi/app/initializer/conf"
	"github.com/atpuxiner/grapi/app/initializer/logger"
	"github.com/atpuxiner/grapi/app/middleware"
	_ "github.com/atpuxiner/grapi/docs"
	"github.com/gin-gonic/gin"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type registerRouterType = func(rg *gin.RouterGroup)

var (
	V1Routers []registerRouterType
)

func InitRouter() *gin.Engine {
	gin.SetMode(conf.Conf.Server.GinMode)
	engine := gin.New()
	engine.Use(middleware.ZapLogger(logger.Logger)).Use(middleware.GinRecovery(logger.Logger, true))
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))
	// ping
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	// v1（此处可添加公共中间件）
	v1 := engine.Group("/api/v1")
	// register
	for _, V1Router := range V1Routers {
		V1Router(v1)
	}
	return engine
}

func registerV1Router(fn registerRouterType) {
	if fn != nil {
		V1Routers = append(V1Routers, fn)
	}
}
