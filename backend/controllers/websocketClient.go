
package controllers

import (
	"net/http"

	"flux-web/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

type client struct {
	ws *websocket.Conn
	send chan []byte
}


var releaseChannel = make(chan models.ReleaseResult)

func (this *WebSocketController) ReleaseWorkloads() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	c := &client{
		send: make(chan []byte),
		ws: ws,
	}

	h.register <- c

	go c.writePump()
}

func (c *client) writePump() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case message := <-c.send:
            if err := c.ws.WriteMessage(websocket.TextMessage, message); err != nil {
                l.Println(err)
                return
            }
		}
	}
}
