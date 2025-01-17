package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("user_list", baseApi.UserList)
		//	配置
		//增
		userRouter.POST("user_config", baseApi.SetUserConfig)
		//删
		//改
		userRouter.PUT("user_config", baseApi.EditUserConfig)
		//查
		userRouter.GET("user_config", baseApi.GetUserConfig)
	}
	return userRouter
}
