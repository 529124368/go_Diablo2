package engine

import (
	"fmt"
	"game/engine/ws"
	"game/role/human"
	"game/status"

	"game/tools"
	"math"
	"runtime"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//切换游戏场景
func (g *Game) ChangeScene(name string) {
	//角色选择场景
	if name == "select" {
		//选择角色场景
		status.Config.CurrentGameScence = tools.GAMESCENESELECTROLE
		g.ui.LoadGameCharaSelectImages()
		runtime.GC()
	} else if name == "game" {
		//进入游戏场景
		status.Config.CurrentGameScence = tools.GAMESCENESTART
		w := sync.WaitGroup{}
		var ww *ws.WsNetManage = nil
		if status.Config.IsNetPlay {
			//网络链接
			ww = ws.NewNet(status.Config)
			g.Ws = ww.Con
			go ww.Start()
		}
		w.Add(3)
		//Palyer Init
		go func() {
			g.player = human.NewPlayer(5280, 1880, tools.IDLE, 0, 0, 0, &asset, g.mapManage, ww, g.ui, g.music)
			g.player.LoadImages("ba", "/man/warrior/", 2)
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
		go func() {
			runtime.GC()
		}()
	}
}

//Draw Game Update
func (g *Game) changeScenceGameUpdate() {
	g.count++
	g.count %= 100
	g.countForMap++
	if !status.Config.MusicIsPlay {
		//音乐
		status.Config.MusicIsPlay = true
		g.music.PlayMusic("Bar_act2_complete_tombs.wav", tools.MUSICWAV)
	}
	if g.player.State != tools.ATTACK && g.player.State != tools.SkILL {
		g.player.State = tools.IDLE
	}
	g.player.MouseX = mouseX
	g.player.MouseY = mouseY

	//鼠标滚轮控制
	if _, x := ebiten.Wheel(); x != 0 {
		status.Config.MapZoom += int(x)
	}
	//主机玩家控制
	g.player.PlayerContr(&g.count)
	//事件循环监听 是否有按钮点击事件
	g.ui.EventLoop(mouseX, mouseY)
}

//Draw Game Scence
func (g *Game) ChangeScenceGameDraw(screen *ebiten.Image) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("has error is x:", r)
		}
	}()
	//获取玩家当前的地图块坐标
	mapX, mapY := tools.GetFloorPositionAt(g.player.X, g.player.Y)
	//玩家所在地图的逻辑坐标
	status.Config.MapTitleX = mapX
	status.Config.MapTitleY = mapY
	g.mapManage.SortLayer(mapX, mapY)
	//Draw floor
	g.mapManage.RenderFloor(screen, status.Config.CamerOffsetX, status.Config.CamerOffsetY)
	//Draw drop items
	g.mapManage.RenderDropItems(screen, status.Config.CamerOffsetX, status.Config.CamerOffsetY, g.player.X, g.player.Y, mouseX, mouseY)
	//切换渲染顺序
	if status.Config.DisplaySort {
		//Draw player
		g.player.Render(screen)
		if status.Config.IsNetPlay && len(g.playerAI) > 0 {
			for _, v := range g.playerAI {
				v.Render(screen)
			}
		}
		//Draw Wall
		g.mapManage.RenderWall(screen, status.Config.CamerOffsetX, status.Config.CamerOffsetY)
		//Draw map Anmi
		g.mapManage.Render(screen, countsFor20, countsFor8, status.Config.CamerOffsetX, status.Config.CamerOffsetY)

	} else {
		//Draw Wall
		g.mapManage.RenderWall(screen, status.Config.CamerOffsetX, status.Config.CamerOffsetY)
		//Draw map Anmi
		g.mapManage.Render(screen, countsFor20, countsFor8, status.Config.CamerOffsetX, status.Config.CamerOffsetY)
		//Draw player
		g.player.Render(screen)
		if status.Config.IsNetPlay && len(g.playerAI) > 0 {
			for _, v := range g.playerAI {
				v.Render(screen)
			}
		}
	}
	//Draw UI
	g.ui.DrawUI(screen)

	//Draw Drop items Anm
	if status.Config.IsPlayDropAnmi {
		if g.mapManage.PlayDropItemAnm(screen, g.player.X, g.player.Y, status.Config.DropItemName, countsFor17) {
			countsFor17 = 0
			status.Config.IsPlayDropAnmi = false
		}
	}
	//Draw Debug
	if status.Config.DisPlayDebugInfo {
		len := tools.Distance(status.Config.PLAYERCENTERX, status.Config.PLAYERCENTERY, int64(mouseX), int64(mouseY))
		re := tools.Angle(math.Abs(float64(int64(mouseY)-status.Config.PLAYERCENTERY)), len)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nplayer world position %d,%d\nmouse position %d,%d\ndir %d\nAngle %f\nCell XY %d,%d",
			int64(ebiten.CurrentFPS()), int64(g.player.X), int64(g.player.Y), g.player.MouseX, g.player.MouseY, tools.CaluteDir(status.Config.PLAYERCENTERX, status.Config.PLAYERCENTERY, int64(g.player.MouseX), int64(g.player.MouseY)), re, mapX, mapY))
	}

	//Change map Frame
	if g.countForMap > 4 {
		countsFor20++
		countsFor20 %= 20
		countsFor8++
		countsFor8 %= 8
		//播放掉落动画
		if status.Config.IsPlayDropAnmi {
			countsFor17++
		}
		g.countForMap = 0
	}
}
