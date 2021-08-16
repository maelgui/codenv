package main

import (
	"codenv-api/controllers"
	"codenv-api/docker"
	"codenv-api/middlewares"
	"codenv-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	docker.ConnectDocker()

	r := gin.Default()
	r.Use(middlewares.ReverseProxyMiddleware())
	api := r.Group("/api")
	{
		api.GET("/workspaces", controllers.ListWorkspace)
		api.POST("/workspaces", controllers.CreateWorkspace)
		api.GET("/workspaces/:id", controllers.RetrieveWorkspace)
		api.DELETE("/workspaces/:id", controllers.DeleteWorkspace)
		api.GET("/workspaces/:id/stop", controllers.StopContainer)
		api.GET("/workspaces/:id/start", controllers.StartContainer)
		api.GET("/workspaces/:id/restart", controllers.RestartContainer)
		api.GET("/workspaces/:id/exec", controllers.Terminal)
	}

	r.Any("/proxy/:id/:port/*path", controllers.Proxy)

	r.Run()
}
