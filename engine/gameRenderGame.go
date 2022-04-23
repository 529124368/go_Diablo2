package engine

import (
	"fmt"
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
			g.player.LoadImages()
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
			g.maps.LoadMap()
			//加载动画
			g.objectA.LoadAnm()
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
	g.countForMap++
	if !g.status.MusicIsPlay {
		//音乐
		g.status.MusicIsPlay = true
		g.music.PlayMusic("Bar_act2_complete_tombs.wav", tools.MUSICWAV)
	}
	if g.player.State != tools.ATTACK {
		g.player.State = tools.IDLE
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.player.MouseX = mouseX
		g.player.MouseY = mouseY
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
		}
	}

	//测试TODO 鼠标滚轮控制
	if _, x := ebiten.Wheel(); x != 0 {
		g.status.MapZoom += int(x)
	}

	//计算鼠标位置
	dir := tools.CaluteDir(g.status.PLAYERCENTERX, g.status.PLAYERCENTERY, int64(g.player.MouseX), int64(g.player.MouseY))

	//TODO 技能攻击
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !g.status.IsTakeItem {
		g.player.SkillName = "liehuo"
		if g.player.State != tools.ATTACK {
			counts = 0
		}
		g.status.Flg = false
		if g.player.Direction != dir || g.player.State != tools.ATTACK {
			g.player.SetPlayerState(tools.ATTACK, dir)

		}
	}
	//事件循环监听 是否有按钮点击事件
	g.ui.EventLoop(mouseX, mouseY)
	//鼠标移动控制
	if !g.status.OpenBag || g.status.OpenBag && mouseX <= tools.LAYOUTX/2 {
		//
		if g.player.OldDirection != g.player.Direction && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			if !g.status.CalculateEnd {
				newPath = tools.CalculateDirPath(g.player.OldDirection, g.player.Direction)
				g.status.CalculateEnd = true
			}
			if len(newPath) >= 3 {
				if turnLoop >= uint8(len(newPath)) {
					turnLoop = uint8(len(newPath) - 1)
					dir = newPath[turnLoop]
					g.player.UpdateOldPlayerDir(g.player.Direction)
				} else {
					dir = newPath[turnLoop]
				}
				turnLoop++
				g.player.SetPlayerState(tools.IDLE, dir)
			} else {
				g.status.CalculateEnd = false
				turnLoop = 0
				g.player.UpdateOldPlayerDir(g.player.Direction)
				g.player.GetMouseController(dir)
			}

		} else {
			g.status.CalculateEnd = false
			turnLoop = 0
			g.player.UpdateOldPlayerDir(g.player.Direction)
			g.player.GetMouseController(dir)
		}

	}
	//根据状态改变帧数
	if g.player.State == tools.IDLE {
		frameNums = 16
		frameSpeed = 5
		g.player.SetPlayerState(tools.IDLE, dir)

	} else if g.player.State == tools.ATTACK {
		frameNums = 16
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
			fmt.Println("has error is :", r)
		}
	}()
	//Draw floor
	g.maps.RenderFloor(screen, g.status.MoveOffsetX, g.status.MoveOffsetY)
	//Draw player
	g.player.Render(screen, counts)
	//Draw Wall
	g.maps.RenderWall(screen, g.status.MoveOffsetX, g.status.MoveOffsetY)
	//Draw map Anmi
	g.objectA.Render(screen, countsForMap, g.status.MoveOffsetX, g.status.MoveOffsetY)
	//Draw UI
	g.ui.DrawUI(screen)
	//获取玩家当前的地图块坐标
	mapX, mapY := tools.GetFloorPositionAt(g.player.X, g.player.Y)
	g.status.MapTitleX = mapX
	g.status.MapTitleY = mapY
	//Draw Debug
	if g.status.DisPlayDebugInfo {
		len := tools.Distance(g.status.PLAYERCENTERX, g.status.PLAYERCENTERY, int64(mouseX), int64(mouseY))
		re := tools.Angle(math.Abs(float64(int64(mouseY)-g.status.PLAYERCENTERY)), len)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nplayer position %d,%d\nmouse position %d,%d\ndir %d\nAngle %f\nmapXY %d,%d",
			int64(ebiten.CurrentFPS()), int64(g.player.X), int64(g.player.Y), g.player.MouseX, g.player.MouseY, tools.CaluteDir(g.status.PLAYERCENTERX, g.status.PLAYERCENTERY, int64(g.player.MouseX), int64(g.player.MouseY)), re, mapX, mapY))
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
	if g.countForMap > 5 {
		countsForMap++
		g.countForMap = 0
		if countsForMap >= 20 {
			countsForMap = 0
		}
	}
}
