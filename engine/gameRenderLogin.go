package engine

import (
	"fmt"
	"game/controller"
	"game/status"
	"game/tools"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Draw Login Update
func (g *Game) ChangeScenceLoginUpdate() {
	//切换场景判定
	if controller.MouseOnceLeftPress() || controller.IsTouch() {
		if mouseX > 286 && mouseX < 503 && mouseY > 150 && mouseY < 218 || controller.Touch(286, 150, 503, 218) {
			status.Config.ChangeScenceFlg = true
			//切换场景
			g.ChangeScene("select")
			status.Config.ChangeScenceFlg = false
		}
	}
	//音乐控制
	if !status.Config.MusicIsPlay {
		status.Config.MusicIsPlay = true
		g.music.PlayBGMusic("Act0-Intro.mp3", tools.MUSICMP3)
	}
	g.count++
	//Change Frame
	if g.count > 2 {
		counts++
		counts %= 30
		g.count = 0
	}
}

//Draw Login Scence
func (g *Game) ChangeScenceLoginDraw(screen *ebiten.Image) {
	//Draw UI
	g.ui.DrawUI(screen, mouseX, mouseY)

	//Draw Logo Left
	name := "logoFireLeft_" + strconv.Itoa(counts)
	left, _, _ := g.ui.GetAnimator(tools.PlistN, name)
	opLo := &ebiten.DrawImageOptions{}
	opLo.Filter = ebiten.FilterLinear
	opLo.GeoM.Translate(220, 0)
	opLo.CompositeMode = ebiten.CompositeModeLighter
	opLo.GeoM.Scale(1, 0.7)
	screen.DrawImage(left, opLo)
	//Draw Logo Right
	name = "logoFireRight_" + strconv.Itoa(counts)
	right, _, _ := g.ui.GetAnimator(tools.PlistN, name)
	opRo := &ebiten.DrawImageOptions{}
	opRo.Filter = ebiten.FilterLinear
	opRo.GeoM.Translate(float64(220+right.Bounds().Size().X), 0)
	opRo.CompositeMode = ebiten.CompositeModeLighter
	opRo.GeoM.Scale(1, 0.7)
	screen.DrawImage(right, opRo)
	if status.Config.DisPlayDebugInfo {
		//Draw Debug
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nmouse position %d,%d\n",
			int64(ebiten.CurrentFPS()), mouseX, mouseY))
	}
	//Change Frame
	if g.count > frameSpeed {
		counts++
		counts %= 30
		g.count = 0
	}
}
