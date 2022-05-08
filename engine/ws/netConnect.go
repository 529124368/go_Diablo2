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
	url := "ws://127.0.0.1:8082/game?ConToken=zimugeWO**@erfs45656DGKZNNSJD"
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
	// time.Sleep(time.Second)
	w.SendMessage("@@whoNotMe")
	//接收消息
	go func() {
		for {
			s, err := w.reciveMessage()
			if err != nil {
				return
			}
			//放入消息队列
			w.status.Queue <- s
		}
	}()
	//心跳维持
	for {
		time.Sleep(time.Second * 500)
		err := w.SendMessage("@@ping")
		if err != nil {
			return
		}
	}
}

//往客户端发送消息
func (w *WsNetManage) SendMessage(msg string) error {
	err := w.Con.WriteMessage(1, []byte(msg))
	return err
}

//接受服务器端消息
func (w *WsNetManage) reciveMessage() ([]byte, error) {
	_, mesg, err := w.Con.ReadMessage()
	return mesg, err
}
