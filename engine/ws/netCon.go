package ws

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WsNetManage struct {
	Con *websocket.Conn
}

func NewNet() *WsNetManage {
	url := "ws://124.220.178.68:8081/game?ConToken=zimugeWO**@erfs45656DGKZNNSJD"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	w := &WsNetManage{
		Con: ws,
	}
	return w
}

func (w *WsNetManage) Start() {
	w.Con.WriteMessage(1, []byte("@@InsertRoom|1"))
	//心跳维持
	go func() {
		for {
			w.Con.WriteMessage(1, []byte("@@ping"))
			time.Sleep(time.Second * 500)
		}
	}()
	//接收消息
	go func() {
		for {
			_, mesg, _ := w.Con.ReadMessage()
			fmt.Println(string(mesg))
		}

	}()
}
