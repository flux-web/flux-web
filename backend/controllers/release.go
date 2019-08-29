
package controllers

import (
	"time"
	"net/http"

	"flux-web/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

type ReleaseResult struct{
	RequestID string
	Status int
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

	go func(ws *websocket.Conn){
		//for {
			for releaseResult := range releaseChannel{
				l.Printf("got new msg in channel: " + releaseResult.Status)
				//if err := ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
				//	l.Printf("error in ws.WriteMessage: ")
				//	l.Println(err)
				//}
				if err := ws.WriteJSON(releaseResult); err != nil{
					l.Printf("error in ws.WriteMessage: ")
					l.Println(err)
				}
				l.Println("msg from ws sent successfully")
			}
		//}
	}(ws)
}

func triggerRelease() []byte{
	time.Sleep(time.Second)
	return []byte("worked!")
}