package controllers

import (
	"codenv-api/docker"
	"codenv-api/models"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Terminal(c *gin.Context) {
	var workspace models.Workspace

	if err := models.DB.Where("id = ?", c.Param("id")).First(&workspace).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	containerConn := docker.ExecContainer(workspace.ContainerID)

	wshandler(c.Writer, c.Request, containerConn)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wshandler(w http.ResponseWriter, r *http.Request, containerConn types.HijackedResponse) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	defer conn.Close()
	defer containerConn.Close()

	// Received of Container Docker
	go func() {
		for {
			buffer := make([]byte, 4096, 4096)
			c, err := containerConn.Reader.Read(buffer)
			if err != nil {
				conn.Close()
				break
			}
			if c > 0 {
				conn.WriteMessage(1, buffer[8:c])
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			containerConn.Close()
			break
		}
		containerConn.Conn.Write(msg)
	}
}
