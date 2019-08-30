
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
		for releaseResult := range releaseChannel{
			l.Printf("got new msg in channel: " + releaseResult.Status)
			err = ws.WriteJSON(releaseResult)
			if err != nil {
				// End request if socket is closed
				if isExpectedClose(err) {
					l.Println("Expected Close on socket", err)
					break
				} else {
					l.Println(err) 
				}
			}else{
				l.Println("msg from ws sent successfully")
			}	
		}
	}(ws)
}

func isExpectedClose(err error) bool {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
		l.Println("Unexpected websocket close: ", err)
		return false
	}

	return true
}