package engine

import (
	"encoding/json"
	"fmt"
	"game/role"
	"game/tools"
	"strconv"
	"strings"
)

var sm serviceMessage

type serviceMessage struct {
	Status bool
	Data   string
	Mes    string
}

//接受服务器端消息
func (g *Game) ListenMessage() {
	for {
		msg := <-g.status.Queue
		//处理消息
		g.Handle(msg)
	}
}

//解包
func (g *Game) unpack(msg []byte) {
	err := json.Unmarshal(msg, &sm)
	if err != nil {
		fmt.Println(err)
	}
}

//新建角色
func (g *Game) CreatePlayer(x, y float64, name, playerName string) {
	r := role.NewPlayer(x, y, tools.IDLE, 0, 0, 0, &asset, g.mapManage, g.status, nil)
	r.PlayerName = playerName
	r.LoadImages(name, 1)
	g.player = append(g.player, r)
}

//删除角色
func (g *Game) DeletePlayer(id int) {
	g.player[id].GC()
	if id < len(g.player)-1 {
		g.player = append(g.player[:id], g.player[id+1:]...)
	} else {
		g.player = g.player[:id]
	}

}

//移动角色
func (g *Game) MovePlayer() {

}

//处理消息
func (g *Game) Handle(msg []byte) {
	//解包
	g.unpack(msg)
	//打印数据
	fmt.Println(sm.Data)
	//动态创建角色 测试用
	if len(sm.Data) > 12 && sm.Data[:12] == "@@newplayer|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 4 {
			d1, _ := strconv.ParseFloat(d[1], 64)
			d2, _ := strconv.ParseFloat(d[2], 64)
			g.CreatePlayer(d1, d2, d[3], "")
			return
		}
	}
	//玩家登陆
	if len(sm.Data) > 10 && sm.Data[:10] == "@@loginIn|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 2 {
			g.CreatePlayer(5202, 1959, "ba", d[1])
			return
		}
	}
	//玩家下线
	if len(sm.Data) > 11 && sm.Data[:11] == "@@loginOut|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 2 {
			for k, v := range g.player {
				if v.PlayerName == d[1] {
					g.DeletePlayer(k)
					return
				}
			}
		}
	}
	//获取名字
	if len(sm.Data) > 7 && sm.Data[:7] == "@@Name|" {
		d := strings.Split(sm.Data, "|")
		g.player[0].PlayerName = d[1]
	}
}
