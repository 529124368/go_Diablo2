package engine

import (
	"encoding/json"
	"fmt"

	"game/role/human"
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
		msg, ok := <-g.status.Queue
		if !ok {
			fmt.Println("afs")
			return
		}
		//处理消息
		g.Handle(msg)
	}
}

//解包
func (g *Game) unpack(msg []byte) {
	err := json.Unmarshal(msg, &sm)
	if err != nil {
		fmt.Println("afss")
		fmt.Println(err)
	}
}

//新建角色
func (g *Game) CreatePlayer(x, y float64, name, playerName string) {
	r := human.NewPlayerAI(x, y, tools.IDLE, 0, g.status, &asset)
	r.PlayerName = playerName
	r.LoadImages(name, 1)
	g.playerAI = append(g.playerAI, r)
}

//删除角色
func (g *Game) DeletePlayer(id int) {
	g.playerAI[id].GC()
	if id < len(g.playerAI)-1 {
		g.playerAI = append(g.playerAI[:id], g.playerAI[id+1:]...)
	} else {
		g.playerAI = g.playerAI[:id]
	}

}

//移动角色
func (g *Game) MovePlayer() {

}

//处理消息
func (g *Game) Handle(msg []byte) {
	//解包
	g.unpack(msg)
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
	//角色移动
	if len(sm.Data) > 7 && sm.Data[:7] == "@@Move|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 5 {
			pm := d[1]
			mx, _ := strconv.ParseFloat(d[2], 64)
			my, _ := strconv.ParseFloat(d[3], 64)
			md, _ := strconv.Atoi(d[4])
			for _, v := range g.playerAI {
				if v.PlayerName == pm {
					v.UpdatePlayerNextMovePositonAI(mx, my, uint8(md))
					return
				}
			}
		}
	}
	//玩家登陆
	if len(sm.Data) > 10 && sm.Data[:10] == "@@loginIn|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 2 {
			g.CreatePlayer(5280, 1880, "ba", d[1])
			return
		}
	}
	//玩家下线
	if len(sm.Data) > 11 && sm.Data[:11] == "@@loginOut|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 2 {
			for k, v := range g.playerAI {
				if v.PlayerName == d[1] {
					g.DeletePlayer(k)
					return
				}
			}
		}
	}
	//玩家停止移动
	if len(sm.Data) > 10 && sm.Data[:10] == "@@MoveEnd|" {
		d := strings.Split(sm.Data, "|")
		if len(d) == 4 {
			for _, v := range g.playerAI {
				if v.PlayerName == d[1] {
					v.StopPlayerMoveAI()
					v.X, _ = strconv.ParseFloat(d[2], 64)
					v.Y, _ = strconv.ParseFloat(d[3], 64)
					return
				}
			}
		}
	}
	//除了我还有谁
	if len(sm.Data) > 12 && sm.Data[:12] == "@@HasPlayer|" {
		d := strings.Split(sm.Data, "|")
		var px float64 = 5280
		var py float64 = 1880
		for _, na := range d[1:] {
			w := strings.Split(na, "%")
			if w[1] != "0" && w[2] != "0" {
				px, _ = strconv.ParseFloat(w[1], 64)
				py, _ = strconv.ParseFloat(w[2], 64)
			}
			g.CreatePlayer(px, py, "ba", w[0])
		}
		return
	}
	//获取名字
	if len(sm.Data) > 7 && sm.Data[:7] == "@@Name|" {
		d := strings.Split(sm.Data, "|")
		g.player.PlayerName = d[1]
		return
	}
}
