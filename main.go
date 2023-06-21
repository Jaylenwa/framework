package main

import (
	"fmt"
	_ "framework/boot"
	"framework/global"
	"framework/interfaces"
	"framework/interfaces/handler"
	"framework/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpUserHandler interfaces.HttpRouterInterface
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

func main() {

	s := &Server{
		httpUserHandler: handler.NewHttpUserHandler(),
	}
	s.Start()

	select {}
}
