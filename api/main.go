package main

import (
	"codenv-api/controllers"
	"codenv-api/docker"
	"codenv-api/middlewares"
	"codenv-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	//models.ConnectDataBase()
	//docker.ConnectDocker()

	r := gin.Default()

	r.Use(middlewares.ReverseProxyMiddleware())
	r.Use(static.Serve("/", static.LocalFile("./static", true)))

	r.Any("/proxy/:id/:port/*path", controllers.Proxy)

	api := r.Group("/api")
	{
		api.GET("/workspaces", controllers.ListWorkspace)
		api.POST("/workspaces", controllers.CreateWorkspace)
		api.GET("/workspaces/:id", controllers.RetrieveWorkspace)
		api.DELETE("/workspaces/:id", controllers.DeleteWorkspace)
		api.GET("/workspaces/:id/stop", controllers.StopContainer)
		api.GET("/workspaces/:id/start", controllers.StartContainer)
		api.GET("/workspaces/:id/restart", controllers.RestartContainer)
		api.GET("/workspaces/:id/exec", controllers.OpenTerminal)

		api.POST("/proxies", controllers.CreateProxy)

	}

	ws := r.Group("/ws")
	{
		ws.GET("/:taskID", controllers.AttachTerminal)
	}

	go killInactiveContainers()

	r.Run()
}

type HealthzResponse struct {
	Status        string `json:"status"`
	LastHeartbeat string `json:"lastHeartbeat"`
}

func killInactiveContainers() {
	for range time.Tick(time.Second * 10) {
		var workspaces []models.Workspace

		models.DB.Preload("Proxies").Find(&workspaces)

		for _, w := range workspaces {
			url, err := controllers.GetProxyRemote(w.ID, "8080")
			if err != nil {
				continue
			}
			fmt.Println("[KILLER] Checking workspace ", w.Name, " (", w.ID, ")")
			url.Path = "/healthz"
			r, err := http.Get(url.String())
			if err != nil {
				continue
			}
			var target HealthzResponse
			json.NewDecoder(r.Body).Decode(&target)

			if target.Status == "expired" {
				docker.StopContainer(w.ContainerID)
				fmt.Println("[KILLER] Stopping workspace ", w.Name, " (", w.ID, ")")
			}
		}
	}
}
