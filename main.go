package main

import (
	"fmt"
	adapterDriver "framework/adapter/driver"
	_ "framework/boot"
	"framework/global"
	"framework/infra/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpUserHandler adapterDriver.HttpRouterInterface
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
		httpUserHandler: adapterDriver.NewHttpUserHandler(),
	}
	s.Start()

	select {}
}
