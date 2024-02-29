package router

import (
	"github.com/atpuxiner/grapi/app/api/v1"
	"github.com/atpuxiner/grapi/app/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter() {
	registerV1Router(func(rg *gin.RouterGroup) {
		userApi := v1.NewUserApi()
		rgg := rg.Group("user") // 此处可添加公共中间件
		{
			rgg.GET("/:id", middleware.JwtAuth(), userApi.Get)
			rgg.GET("", middleware.JwtAuth(), userApi.GetList)
			rgg.POST("", userApi.Create)
			rgg.PUT("/:id", middleware.JwtAuth(), userApi.Update)
			rgg.DELETE("/:id", middleware.JwtAuth(), userApi.Delete)
		}
	})
}

func init() {
	userRouter()
}
