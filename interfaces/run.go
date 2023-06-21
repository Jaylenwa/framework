package interfaces

import (
	"fmt"
	"framework/global"
	"framework/interfaces/handler"
	"framework/interfaces/middleware"
	"framework/interfaces/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpUserHandler router.HttpRouterInterface
}

func (s *Server) Start() {
	go func() {
		engine := gin.New()
		// 注册路由 - 健康检查
		routerHealth := engine.Group("/api/v1")
		routerHealth.Use(middleware.ErrorHandlerMiddleware())
		s.httpUserHandler.RegisterRouterPublic(routerHealth)
		url := fmt.Sprintf("%s:%d", global.GConfig.Project.Host, global.GConfig.Project.Port)
		_ = engine.Run(url)
	}()
}

// 启动项目
func init() {
	s := &Server{
		httpUserHandler: handler.NewHttpUserHandler(),
	}
	s.Start()
	select {}
}
