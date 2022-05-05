package engine

import (
	"fmt"
	"time"
)

//处理服务器端消息
func (g *Game) handle() {
	for {
		msg := <-g.status.Queue
		if msg == "@@newplayer" {
			fmt.Println("创建角色")
		}
		time.Sleep(time.Microsecond)
	}
}
