package ws

import (
	"game/engine/ws/pb"
	"game/status"
	"game/tools"
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
	w.SendMessage(true, "@@myName", "", "", nil)
	w.SendMessage(true, "@@whoNotMe", "", "", nil)
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
		err := w.SendMessage(true, "@@ping", "", "", nil)
		if err != nil {
			return
		}
	}
}

// 往服务器发送消息
func (w *WsNetManage) SendMessage(s bool, f, datas, msg string, p *pb.Player) error {
	err := w.Con.WriteMessage(1, tools.Pack(s, f, datas, msg, p))
	return err
}

// 接受服务器端消息
func (w *WsNetManage) reciveMessage() ([]byte, error) {
	_, mesg, err := w.Con.ReadMessage()
	return mesg, err
}
