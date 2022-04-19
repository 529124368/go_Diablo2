package engine

import (
	"embed"
	"game/fonts"
	"game/layout"
	"game/mapCreator/anm"
	"game/maps"
	"game/music"
	"game/role"
	"game/status"
	"game/tools"
	"runtime"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

//config
const (
	SCREENWIDTH  int = 490
	SCREENHEIGHT int = 300
	WEOFFSETX    int = 127
	WEOFFSETY    int = 14
)

type Game struct {
	count       int
	countForMap int
	player      *role.Player         //玩家
	maps        *maps.MapBase        //地图
	objectA     *anm.Anm             //object 动画
	ui          *layout.UI           //UI
	music       music.MusicInterface //音乐
	status      *status.StatusManage //状态管理器
	font_style  *fonts.FontBase      //字体
}

var (
	counts       int = 0
	countsForMap int = 0
	frameNums    int = 4
	frameSpeed   int = 5
	mouseX       int
	mouseY       int
	newPath      []uint8
	turnLoop     uint8 = 0
)

//GameEngine
func NewGame(asset *embed.FS) *Game {
	//statueManage
	sta := status.NewStatusManage()
	//Map
	m := maps.NewMap(asset, sta)
	//Player  设置初始状态和坐标
	r := role.NewPlayer(5280, 1840, tools.IDLE, 0, 0, 0, asset, m, sta)
	//字体
	f := fonts.NewFont(asset)
	//UI
	u := layout.NewUI(asset, sta, m, f)
	//BGM
	bgm := music.NewMusicBGM(asset)
	//场景动画
	object := anm.NewAnm(asset, sta)

	gameEngine := &Game{
		count:       0,
		countForMap: 0,
		player:      r,
		maps:        m,
		ui:          u,
		music:       bgm,
		status:      sta,
		objectA:     object,
		font_style:  f,
	}
	return gameEngine
}

//引擎启动
func (g *Game) StartEngine() {
	//隐藏鼠标系统的ICON
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	w := sync.WaitGroup{}
	w.Add(1)
	//UI Init
	go func() {
		g.font_style.LoadFont("resource/font/DiabloHeavy.ttf")
		g.ui.LoadGameLoginImages()
		runtime.GC()
		w.Done()
	}()
	w.Wait()
	go func() {
		runtime.GC()
	}()
}

func (g *Game) Update() error {
	mouseX, mouseY = ebiten.CursorPosition()
	//切换场景逻辑
	if !g.status.ChangeScenceFlg {
		if g.status.CurrentGameScence == tools.GAMESCENESTART {
			//进入游戏场景逻辑
			g.changeScenceGameUpdate()
		} else if g.status.CurrentGameScence == tools.GAMESCENEOPENDOOR {
			//游戏加载逻辑
			g.ChangeScenceOpenDoorUpdate()
		} else if g.status.CurrentGameScence == tools.GAMESCENESELECTROLE {
			//进入游戏选择界面逻辑
			g.ChangeScenceSelectUpdate()
		} else {
			//进入游戏登录界面逻辑
			g.ChangeScenceLoginUpdate()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//判断是否切换场景
	if !g.status.ChangeScenceFlg {
		if g.status.CurrentGameScence == tools.GAMESCENESTART {
			g.ChangeScenceGameDraw(screen)
		} else if g.status.CurrentGameScence == tools.GAMESCENESELECTROLE {
			g.ChangeScenceSelectDraw(screen)
		} else if g.status.CurrentGameScence == tools.GAMESCENEOPENDOOR {
			g.ChangeScenceOpenDoorDraw(screen)
		} else {
			g.ChangeScenceLoginDraw(screen)
		}
	}
	//绘制鼠标ICON
	g.ui.DrawMouseIcon(screen, mouseX, mouseY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tools.LAYOUTX, tools.LAYOUTY
}
