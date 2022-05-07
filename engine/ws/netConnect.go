package ws

import (
	"game/status"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WsNetManage struct {
	Con    *websocket.Conn
	status *status.StatusManage
}

func NewNet(s *status.StatusManage) *WsNetManage {
	url := "ws://124.220.178.68:8082/game?ConToken=zimugeWO**@erfs45656DGKZNNSJD"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	w := &WsNetManage{
		Con:    ws,
		status: s,
	}
	return w
}

func (w *WsNetManage) Start() {
	//发送消息
	w.SendMessage("@@myName")
	//心跳维持
	go func() {
		for {
			w.SendMessage("@@ping")
			time.Sleep(time.Second * 500)
		}
	}()
	//接收消息
	go func() {
		for {
			s := w.reciveMessage()
			//放入消息队列
			w.status.Queue <- s
		}
	}()
}

//往客户端发送消息
func (w *WsNetManage) SendMessage(msg string) {
	w.Con.WriteMessage(1, []byte(msg))
}

//接受服务器端消息
func (w *WsNetManage) reciveMessage() []byte {
	_, mesg, _ := w.Con.ReadMessage()
	return mesg
}
