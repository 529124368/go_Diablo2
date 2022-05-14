package engine

import (
	"fmt"
	"game/role/human"
	"game/status"
	"game/tools"
	"strconv"
	"strings"
)

//接受服务器端消息
func (g *Game) ListenMessage() {
	for {
		msg, ok := <-status.Config.Queue
		if !ok {
			return
		}
		//处理消息
		g.Handle(msg)
	}
}

//新建角色
func (g *Game) CreatePlayer(x, y float64, name, playerName string) {
	r := human.NewPlayerAI(x, y, tools.IDLE, 0, &asset)
	r.PlayerName = playerName
	r.LoadImages(name, "/man/warrior/", 2)
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

//处理消息
func (g *Game) Handle(msg []byte) {
	//解包
	sm := tools.Unpack(msg)
	//角色移动
	if sm.Flag == "@@Move" {
		for _, v := range g.playerAI {
			fmt.Println(sm)
			if v.PlayerName == sm.Data.Man.Name {
				v.UpdatePlayerPositonAI(sm.Data.Man.X, sm.Data.Man.Y, uint8(sm.Data.Man.Dir), sm.Data.Man.State)
				return
			}
		}
	}
	//玩家登陆
	if sm.Flag == "@@loginIn" {
		g.CreatePlayer(5280, 1880, "ba", sm.Data.Data)
		return
	}
	//玩家下线
	if sm.Flag == "@@loginOut" {
		for k, v := range g.playerAI {
			if v.PlayerName == sm.Data.Data {
				g.DeletePlayer(k)
				return
			}
		}
	}
	//玩家停止移动
	if sm.Flag == "@@MoveEnd" {
		for _, v := range g.playerAI {
			if v.PlayerName == sm.Data.Man.Name {
				v.X = sm.Data.Man.X
				v.Y = sm.Data.Man.Y
				v.State = tools.IDLE
				return
			}
		}
	}
	//除了我还有谁
	if sm.Flag == "@@HasPlayer" {
		d := strings.Split(sm.Data.Data, "|")
		var px float64 = 5280
		var py float64 = 1880
		for _, na := range d {
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
	if sm.Flag == "@@Name" {
		g.player.PlayerName = sm.Data.Data
		return
	}
}
