package Websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ?连接对象全局定义
var Conn *websocket.Conn

func Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()
	Conn = conn
}

func SendClientMsg(msg string) {
	Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}
