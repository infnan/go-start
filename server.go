package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func testRoute(c *gin.Context) {
	c.String(200, "Hello, world!")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  20480,
	WriteBufferSize: 20480,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func testWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
	}

	defer conn.Close()

	for {
		if conn == nil {
			break
		}
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println(message)
	}
}

func startServer() {
	// 接口服务器
	gin.SetMode(gin.ReleaseMode)
	serv := gin.New()

	group := serv.Group("/")
	g1 := group.Group("/")
	g1.GET("/test", testRoute)
	g1.GET("/ws", testWs)

	log.Info("Listening...")
	serv.Run("0.0.0.0:9000")
}
