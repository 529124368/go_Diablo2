package engine

import (
	"fmt"
	"game/controller"
	"game/engine/ws"
	"game/tools"
	"math"
	"runtime"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var dropItemName = ""
var newpositonX, newpositonY float64
var newDir uint8

//切换游戏场景
func (g *Game) ChangeScene(name string) {
	//角色选择场景
	if name == "select" {
		//选择角色场景
		g.status.CurrentGameScence = tools.GAMESCENESELECTROLE
		g.ui.LoadGameCharaSelectImages()
		runtime.GC()
	} else if name == "game" {
		//进入游戏场景
		g.status.CurrentGameScence = tools.GAMESCENESTART
		w := sync.WaitGroup{}
		w.Add(3)
		//Palyer Init
		go func() {
			g.player[0].LoadImages("ba2", 2)
			runtime.GC()
			w.Done()
		}()
		//Ui
		go func() {
			g.ui.LoadGameImages()
			runtime.GC()
			w.Done()
		}()
		//Map Init
		go func() {
			g.mapManage.LoadMap()
			//加载动画
			g.mapManage.LoadAnm()
			runtime.GC()
			w.Done()
		}()
		w.Wait()
		//网络链接
		ww := ws.NewNet(g.status)
		ww.Start()
		go func() {
			runtime.GC()
		}()
	}
}

//Draw Game Update
func (g *Game) changeScenceGameUpdate() {
	g.count++
	g.countForMap++
	if !g.status.MusicIsPlay {
		//音乐
		g.status.MusicIsPlay = true
		g.music.PlayMusic("Bar_act2_complete_tombs.wav", tools.MUSICWAV)
	}
	//背景音乐是否还在播放
	// if !g.music.IsPlayingBGMusic() {
	// 	go func() {
	// 		g.music.PlayBGMusic("town1.wav", tools.MUSICWAV)
	// 	}()
	// }
	if g.player[0].State != tools.ATTACK && g.player[0].State != tools.SkILL {
		g.player[0].State = tools.IDLE
	}
	g.player[0].MouseX = mouseX
	g.player[0].MouseY = mouseY
	//计算鼠标位置
	dir := tools.CaluteDir(g.status.PLAYERCENTERX, g.status.PLAYERCENTERY, int64(g.player[0].MouseX), int64(g.player[0].MouseY))

	//鼠标事件
	if controller.MouseleftPress() || controller.IsTouch() {
		//防止点击UI界面也移动
		if mouseY < 436 {
			g.status.Flg = true
		}
		//如果打开包裹，包裹已右位置不能点击移动
		if g.status.OpenBag && mouseX >= tools.LAYOUTX/2 {
			g.status.Flg = false
		}
		//如果打开MINi板子，并且没有打开包裹 以下坐标不可以点击移动
		if g.status.OpenMiniPanel && !g.status.OpenBag && mouseX >= 305 && mouseX <= 475 && mouseY > 407 {
			g.status.Flg = false
		}
		//如果打开MINi板子，并且打开包裹 以下坐标不可以点击移动
		if g.status.OpenMiniPanel && g.status.OpenBag && mouseX >= 205 && mouseX <= 377 && mouseY > 407 {
			g.status.Flg = false
		}
		//如果拿起物品也不可以移动
		if g.status.IsTakeItem {
			g.status.Flg = false
			//这个范围内就是丢弃物品
			if mouseY < 436 && mouseX < 390 {
				if !g.status.IsDropDeal {
					g.status.IsDropDeal = true
					//播放掉落物品动画
					g.status.IsPlayDropAnmi = true
					//音乐
					g.music.PlayMusic("diaoluo.mp3", tools.MUSICMP3)
					//丢弃物品
					dropItemName = g.ui.ClearTempBag()
				}
			}
		}
		if g.status.Flg {
			//计算新的位置
			newDir = dir
			newpositonX = g.player[0].X + float64(mouseX) - 395
			newpositonY = g.player[0].Y + float64(mouseY) - 240
		}
	}

	if newpositonX != 0 && newpositonY != 0 && (math.Abs(g.player[0].X-newpositonX) > 1 && math.Abs(g.player[0].Y-newpositonY) > 1) {
		g.status.Flg = true
		//鼠标人物移动控制
		g.player[0].PlayerMove(&newDir)
	} else {
		newpositonX = 0
		newpositonY = 0
	}

	//鼠标滚轮控制
	if _, x := ebiten.Wheel(); x != 0 {
		g.status.MapZoom += int(x)
	}

	//普通攻击
	if controller.MouseRightPress() && !g.status.IsTakeItem {
		//g.player[0].WsCon.SendMessage("attack#######")
		if g.player[0].State != tools.ATTACK {
			counts = 0
		}
		g.status.Flg = false
		if g.player[0].Direction != dir || g.player[0].State != tools.ATTACK {
			g.player[0].SetPlayerState(tools.ATTACK, dir)
		}
	}
	//技能
	if controller.MousePressF1() && !g.status.IsTakeItem {
		//音乐
		g.music.PlayMusic("File00002184.wav", tools.MUSICWAV)
		//g.player[0].SkillName = "狂风"
		if g.player[0].State != tools.SkILL {
			counts = 0
		}
		g.status.Flg = false
		if g.player[0].Direction != dir || g.player[0].State != tools.SkILL {
			g.player[0].SetPlayerState(tools.SkILL, dir)
		}
	}
	//事件循环监听 是否有按钮点击事件
	g.ui.EventLoop(mouseX, mouseY)

	//根据状态改变帧数
	if g.player[0].State == tools.IDLE {
		frameNums = 16
		frameSpeed = 5

	} else if g.player[0].State == tools.ATTACK {
		frameNums = 16
		frameSpeed = 1
	} else if g.player[0].State == tools.SkILL {
		frameNums = 14
		frameSpeed = 1
	} else {
		frameNums = 8
		frameSpeed = 5
	}
}

//Draw Game Scence
func (g *Game) ChangeScenceGameDraw(screen *ebiten.Image) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("has error is x:", r)
		}
	}()
	//获取玩家当前的地图块坐标
	mapX, mapY := tools.GetFloorPositionAt(g.player[0].X, g.player[0].Y)
	//玩家所在地图的逻辑坐标
	g.status.MapTitleX = mapX
	g.status.MapTitleY = mapY
	g.mapManage.SortLayer(mapX, mapY)
	//Draw floor
	g.mapManage.RenderFloor(screen, g.status.CamerOffsetX, g.status.CamerOffsetY)
	//Draw drop items
	g.mapManage.RenderDropItems(screen, g.status.CamerOffsetX, g.status.CamerOffsetY, g.player[0].X, g.player[0].Y)
	//切换渲染顺序
	if g.status.DisplaySort {
		//Draw player
		for k, v := range g.player {
			if k == 0 {
				v.Render(screen, counts)
			} else {
				v.RenderCopy(screen, counts)
			}
		}
		//Draw Wall
		g.mapManage.RenderWall(screen, g.status.CamerOffsetX, g.status.CamerOffsetY)
		//Draw map Anmi
		g.mapManage.Render(screen, countsFor20, countsFor8, g.status.CamerOffsetX, g.status.CamerOffsetY)

	} else {
		//Draw Wall
		g.mapManage.RenderWall(screen, g.status.CamerOffsetX, g.status.CamerOffsetY)
		//Draw map Anmi
		g.mapManage.Render(screen, countsFor20, countsFor8, g.status.CamerOffsetX, g.status.CamerOffsetY)
		//Draw player
		for k, v := range g.player {
			if k == 0 {
				v.Render(screen, counts)
			} else {
				v.RenderCopy(screen, counts)
			}
		}
	}
	//Draw UI
	g.ui.DrawUI(screen)

	//Draw Drop items Anm
	if g.status.IsPlayDropAnmi {
		if g.mapManage.PlayDropItemAnm(screen, g.player[0].X, g.player[0].Y, dropItemName, countsFor17) {
			countsFor17 = 0
			g.status.IsPlayDropAnmi = false
		}
	}
	//Draw Debug
	if g.status.DisPlayDebugInfo {
		len := tools.Distance(g.status.PLAYERCENTERX, g.status.PLAYERCENTERY, int64(mouseX), int64(mouseY))
		re := tools.Angle(math.Abs(float64(int64(mouseY)-g.status.PLAYERCENTERY)), len)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nplayer world position %d,%d\nmouse position %d,%d\ndir %d\nAngle %f\nCell XY %d,%d",
			int64(ebiten.CurrentFPS()), int64(g.player[0].X), int64(g.player[0].Y), g.player[0].MouseX, g.player[0].MouseY, tools.CaluteDir(g.status.PLAYERCENTERX, g.status.PLAYERCENTERY, int64(g.player[0].MouseX), int64(g.player[0].MouseY)), re, mapX, mapY))
	}
	//Change player Frame
	if g.count > frameSpeed {
		counts++
		g.count = 0
		if counts >= frameNums {
			counts = 0
		}
	}

	//Change map Frame
	if g.countForMap > 4 {
		countsFor20++
		countsFor8++
		//播放掉落动画
		if g.status.IsPlayDropAnmi {
			countsFor17++
		}
		g.countForMap = 0
		if countsFor20 >= 20 {
			countsFor20 = 0
		}
		if countsFor8 >= 8 {
			countsFor8 = 0
		}

	}
}
