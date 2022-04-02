package main

import (
	"embed"
	"fmt"
	"game/layout"
	"game/maps"
	"game/role"
	"game/tools"
	_ "image/png"
	"log"
	"math"
	"runtime"
	"strconv"
	"sync"

	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//config
const (
	SCREENWIDTH   int   = 490
	SCREENHEIGHT  int   = 300
	OFFSETX       int   = -30
	OFFSETY       int   = -30
	LAYOUTX       int   = 790
	LAYOUTY       int   = 480
	PLAYERCENTERX int64 = 388 //LAYOUTX/2
	PLAYERCENTERY int64 = 242 //LAYOUTY/2
	WEOFFSETX     int   = 127
	WEOFFSETY     int   = 14
)

var game *Game

//Game
type Game struct {
	count  int
	player *role.Player
	maps   *maps.MapBase
	ui     *layout.UI
	//monster *role.Monster
}

var (
	counts     int  = 0
	frameNums  int  = 4
	flg        bool = false
	frameSpeed int  = 5
)
var op, opS, opMouse, opWea, opSkill *ebiten.DrawImageOptions

//go:embed resource
var images embed.FS

var mouseX, mouseY int

var mouseIcon *ebiten.Image

//factory
func NewGame() *Game {
	//map
	m := maps.NewMap(&images)
	//palayer
	r := role.NewPlayer(float64(LAYOUTX/2), float64(LAYOUTY/2), tools.IDLE, 0, 0, 0, &images, m)
	//UI
	u := layout.NewUI(&images)

	gameStart := &Game{
		count:  0,
		player: r,
		maps:   m,
		ui:     u,
	}
	return gameStart
}

func init() {
	//hidden mouse icon
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	opMouse = &ebiten.DrawImageOptions{}
	s, _ := images.ReadFile("resource/UI/mouse.png")
	mouseIcon = tools.GetEbitenImage(s)
	opS = &ebiten.DrawImageOptions{}
	op = &ebiten.DrawImageOptions{}
	w := sync.WaitGroup{}
	w.Add(3)
	//game init
	game = NewGame()
	//palyer init
	go func() {
		game.player.LoadImages()
		runtime.GC()
		w.Done()
	}()
	//UI init
	go func() {
		game.ui.LoadImages()
		w.Done()
	}()
	//map init
	go func() {
		game.maps.LoadMap()
		w.Done()
	}()
	w.Wait()
	go func() {
		runtime.GC()
	}()
}

func (g *Game) Update() error {
	g.count++
	mouseX, mouseY = ebiten.CursorPosition()
	if g.player.State != tools.ATTACK {
		g.player.State = tools.IDLE
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.player.MouseX = mouseX
		g.player.MouseY = mouseY
		flg = true
	}
	//Calculate direction
	dir := tools.CaluteDir(PLAYERCENTERX, PLAYERCENTERY, int64(g.player.MouseX), int64(g.player.MouseY))
	//EVENT Listen
	g.ui.EventLoop()
	//attack
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		//x, y := ebiten.CursorPosition()
		// //liehuo
		//if x > 583 && x < 627 && y > 380 && y < 424 {
		g.player.SkillName = "liehuo"
		if g.player.State != tools.ATTACK {
			counts = 0
		}
		flg = false
		if g.player.Direction != dir || g.player.State != tools.ATTACK {
			g.player.SetPlayerState(tools.ATTACK, dir)

		}
		//}
		// if x > 621 && x < 664 && y > 331 && y < 373 {
		// 	g.player.SkillName = "chisha"
		// 	if g.player.State != ATTACK {
		// 		counts = 0
		// 	}
		// 	flg = false
		// 	if g.player.Direction != dir || g.player.State != ATTACK {
		// 		g.player.SetPlayerState(ATTACK, dir)

		// 	}
		// }
	}

	//mouse controll
	flg = g.player.GetMouseController(dir, flg)
	//states
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
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("has error is :", r)
		}
	}()
	//draw background
	screen.DrawImage(g.maps.BgImage, g.maps.OpBg)
	//
	name := ""
	//nameSkill := ""
	switch g.player.State {
	case tools.ATTACK:
		name = strconv.Itoa(g.player.Direction) + "_attack_" + strconv.Itoa(counts) + ".png"
		//nameSkill = strconv.Itoa(g.player.Direction) + "_skill_" + strconv.Itoa(counts) + ".png"
	case tools.IDLE:
		name = strconv.Itoa(g.player.Direction) + "_stand_" + strconv.Itoa(counts) + ".png"
	default:
		if counts >= 8 {
			counts = 0
		}
		name = strconv.Itoa(g.player.Direction) + "_run_" + strconv.Itoa(counts) + ".png"
	}
	imagess, x, y := g.player.GetAnimator("man", name)
	//idel -> walk offset
	if g.player.State == tools.RUN {
		x += -4
		y += -3
	}

	//idel -> walk -> attack offset
	if g.player.State == tools.ATTACK {
		x += -50
		y += -30
	}
	//draw shadow
	opS.GeoM.Reset()
	opS.GeoM.Translate(float64(LAYOUTX/2+x-350), float64(LAYOUTY/2+y+365))
	opS.Filter = ebiten.FilterLinear
	opS.GeoM.Rotate(-0.5)
	opS.GeoM.Scale(1, 0.5)
	opS.ColorM.Scale(0, 0, 0, 1)
	opS.ColorM.Translate(0, 0, 0, 0)
	screen.DrawImage(imagess, opS)
	//draw player
	op.GeoM.Reset()
	op.GeoM.Translate(float64(LAYOUTX/2+OFFSETX+x), float64(LAYOUTY/2+OFFSETY+y))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(imagess, op)

	//draw UI
	g.ui.DrawUI(screen)

	//draw mouse icon
	opMouse.GeoM.Reset()
	opMouse.GeoM.Rotate(-0.5)
	opMouse.Filter = ebiten.FilterLinear
	opMouse.GeoM.Translate(float64(mouseX), float64(mouseY))
	screen.DrawImage(mouseIcon, opMouse)

	//draw skill
	// if g.player.State == ATTACK {
	// 	imagey, x, y := g.player.GetAnimator("skill", nameSkill)
	// 	//skill option
	// 	opSkill = &ebiten.DrawImageOptions{}
	// 	opSkill.GeoM.Translate(float64(SCREENWIDTH/2+x), float64(SCREENHEIGHT/2+y))
	// 	opSkill.CompositeMode = ebiten.CompositeModeLighter
	// 	opSkill.GeoM.Scale(1.5, 1.5)
	// 	opSkill.Filter = ebiten.FilterLinear
	// 	screen.DrawImage(imagey, opSkill)
	// }
	//draw debug
	len := tools.Distance(PLAYERCENTERX, PLAYERCENTERY, int64(mouseX), int64(mouseY))
	re := tools.Angle(math.Abs(float64(int64(mouseY)-PLAYERCENTERY)), len)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS %d\nplayer position %d,%d\nmouse position %d,%d\ndir %d\nAngle %f",
		int64(ebiten.CurrentFPS()), int64(g.player.X), int64(g.player.Y), g.player.MouseX, g.player.MouseY, tools.CaluteDir(PLAYERCENTERX, PLAYERCENTERY, int64(g.player.MouseX), int64(g.player.MouseY)), re))
	//change frame
	if g.count > frameSpeed {
		counts++
		g.count = 0
		if counts >= frameNums {
			counts = 0
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return LAYOUTX, LAYOUTY
}

func main() {
	ebiten.SetWindowSize(SCREENWIDTH*2, SCREENHEIGHT*2)
	ebiten.SetWindowTitle("golang Dibalo2")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
