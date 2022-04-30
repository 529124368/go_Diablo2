package engine

import (
	"fmt"
	"game/tools"
	"runtime"
	"strconv"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var clearFlg = false

//Draw OpenDoor Scence
func (g *Game) ChangeScenceOpenDoorDraw(screen *ebiten.Image) {
	if !g.status.DoorCountFlg {
		counts = 0
		g.status.DoorCountFlg = true
	}
	//Draw Open Door
	name := "loading_" + strconv.Itoa(counts)
	door, _, _ := g.ui.GetAnimator(tools.PlistN, name)
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Translate(268, 120)
	screen.DrawImage(door, op)
	//Draw Debug
	if g.status.DisPlayDebugInfo {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nmouse position %d,%d", int64(ebiten.CurrentFPS()), mouseX, mouseY))
	}
}

//Draw OpenDoor Update
func (g *Game) ChangeScenceOpenDoorUpdate() {
	if !clearFlg {
		clearFlg = true
		go func() {
			g.ui.GCSelectBGImage()
			g.ui.ClearSlice(0)
			g.music.CloseBGMusic()
			runtime.GC()
		}()
	}
	g.count++
	//Change Frame
	if g.count > 10 && counts != 9 {
		counts++
		g.count = 0
	}

	// Change Scence
	if counts == 9 && !g.status.LoadingFlg {
		g.status.LoadingFlg = true
		w := sync.WaitGroup{}
		w.Add(1)
		go func() {
			//close music
			g.status.MusicIsPlay = false
			g.ChangeScene("game")
			runtime.GC()
			w.Done()
		}()
		w.Wait()
		g.status.ChangeScenceFlg = false
	}
}
