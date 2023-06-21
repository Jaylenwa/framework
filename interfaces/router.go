package interfaces

import "github.com/gin-gonic/gin"

// HttpRouterInterface 路由公共接口
type HttpRouterInterface interface {
	// RegisterRouterPublic 注册外部API
	RegisterRouterPublic(engine *gin.RouterGroup)

	// RegisterRouterPrivate 注册内部API
	RegisterRouterPrivate(engine *gin.RouterGroup)
}
