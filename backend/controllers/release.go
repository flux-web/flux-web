
package controllers

import (
	"net/http"
	"time"

	//"flux-web/models"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/httplib"
	//"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

type ReleaseResult struct{
	RequestID string
	Status int
}

func (this *WebSocketController) ReleaseWorkloads() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	go func(ws *websocket.Conn){
		for {
			msgType, releaseRequest, err := ws.ReadMessage()
			if err != nil {
				return
			}
			l.Println(string(releaseRequest))

			if err := ws.WriteMessage(msgType, triggerRelease()); err != nil{
				l.Println(err)
				return
			}
		}
	}(ws)
}

func triggerRelease() []byte{
	time.Sleep(time.Second)
	return []byte("worked!")
}