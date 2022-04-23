package engine

import (
	"fmt"
	"game/tools"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Draw Select Update
func (g *Game) ChangeScenceSelectUpdate() {
	//鼠标点击监听
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if mouseX > 622 && mouseX < 702 && mouseY > 428 && mouseY < 456 {
			//获取按钮的回调函数
			f := g.ui.Compents[1].CallFunc()
			//按钮内注册了切换场景
			f(g.ui.Compents[1])
		}
	}
	//音乐播放控制
	if !g.status.MusicIsPlay {
		g.status.MusicIsPlay = true
		g.music.PlayMusic("Act0-Intro.mp3", tools.MUSICMP3)
	}
	//轮图控制
	g.count++
	//Change Frame
	if g.count > 5 {
		counts++
		g.count = 0
		if counts >= 30 {
			counts = 0
		}
	}
}

//Draw Select Scence
func (g *Game) ChangeScenceSelectDraw(screen *ebiten.Image) {
	//Draw UI
	g.ui.DrawUI(screen)
	//Draw Fire 场景火堆
	name := "fire_" + strconv.Itoa(counts) + ".png"
	fire, _, _ := g.ui.GetAnimator("logo", name)
	opf := &ebiten.DrawImageOptions{}
	opf.Filter = ebiten.FilterLinear
	opf.GeoM.Translate(350, 375)
	opf.CompositeMode = ebiten.CompositeModeLighter
	opf.GeoM.Scale(1, 0.7)
	screen.DrawImage(fire, opf)

	//Draw Role 野蛮人战士
	name = "ba_" + strconv.Itoa(counts) + ".png"
	ba, _, _ := g.ui.GetAnimator("role", name)
	opBa := &ebiten.DrawImageOptions{}
	opBa.Filter = ebiten.FilterLinear
	opBa.GeoM.Translate(356, 120)
	screen.DrawImage(ba, opBa)

	//Draw Debug
	if g.status.DisPlayDebugInfo {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nmouse position %d,%d", int64(ebiten.CurrentFPS()), mouseX, mouseY))
	}
	//Draw Text
	g.font_style.Render(screen, 635, 446, "开始游戏", 8, 130, color.White)
}
