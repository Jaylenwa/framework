package router

import "github.com/gin-gonic/gin"

// HttpRouterInterface .
type HttpRouterInterface interface {
	// RegisterRouterPublic 注册外部API
	RegisterRouterPublic(engine *gin.RouterGroup)

	// RegisterRouterPrivate 注册内部API
	RegisterRouterPrivate(engine *gin.RouterGroup)
}
