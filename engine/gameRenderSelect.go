package engine

import (
	"fmt"
	"game/controller"
	"game/status"
	"game/tools"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Draw Select Update
func (g *Game) ChangeScenceSelectUpdate() {
	//鼠标点击监听
	if controller.MouseOnceLeftPress() || controller.IsTouch() {
		if mouseX > 622 && mouseX < 702 && mouseY > 428 && mouseY < 456 || controller.Touch(622, 428, 702, 456) {
			//获取按钮的回调函数
			f := g.ui.Compents[0].CallFunc()
			//按钮内注册了切换场景
			f(g.ui.Compents[0])
		}
	}
	//轮图控制
	g.count++
	//Change Frame
	if g.count > 5 {
		counts++
		g.count = 0
		counts %= 25
	}
}

//Draw Select Scence
func (g *Game) ChangeScenceSelectDraw(screen *ebiten.Image) {
	//背景图
	opm := new(ebiten.DrawImageOptions)
	opm.GeoM.Translate(0, 0)
	opm.GeoM.Scale(1, 0.8)
	screen.DrawImage(g.ui.GetSelectBGImage(), opm)
	//Draw UI
	g.ui.DrawUI(screen)
	//Draw Fire 场景火堆
	name := "fire_" + strconv.Itoa(counts)
	fire, _, _ := g.ui.GetAnimator(tools.PlistN, name)
	if fire != nil {
		opf := &ebiten.DrawImageOptions{}
		opf.Filter = ebiten.FilterLinear
		opf.GeoM.Translate(350, 375)
		opf.CompositeMode = ebiten.CompositeModeLighter
		opf.GeoM.Scale(1, 0.7)
		screen.DrawImage(fire, opf)
	}

	//Draw Role 野蛮人战士
	name = "ba_" + strconv.Itoa(counts)
	ba, _, _ := g.ui.GetAnimator(tools.PlistR, name)
	opBa := &ebiten.DrawImageOptions{}
	opBa.Filter = ebiten.FilterLinear
	opBa.GeoM.Translate(356, 120)
	screen.DrawImage(ba, opBa)

	//Draw Debug
	if status.Config.DisPlayDebugInfo {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nmouse position %d,%d", int64(ebiten.CurrentFPS()), mouseX, mouseY))
	}
	//Draw Text
	g.font_style.Render(screen, 640, 446, "开始游戏", 8, 130, color.White)
}
