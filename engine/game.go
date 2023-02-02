package engine

import (
	"embed"
	"game/controller"
	"game/fonts"
	"game/interfaces"
	"game/layout"
	"game/mapCreator/mapManage"
	"game/music"
	"game/role/human"
	"game/status"
	"game/storage"
	"game/tools"
	"runtime"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 配置信息
const (
	SCREENWIDTH  int = 490
	SCREENHEIGHT int = 300
	WEOFFSETX    int = 127
	WEOFFSETY    int = 14
)

type Game struct {
	count, countForMap int
	player             *human.Player             //本机玩家
	playerAI           []*human.PlayerAI         //AI玩家
	mapManage          interfaces.MapInterface   //地图等管理
	ui                 *layout.UI                //UI
	music              interfaces.MusicInterface //音乐
	font_style         *fonts.FontBase           //字体
	Ws                 *websocket.Conn
}

var (
	counts      int = 0
	countsFor20 int = 0
	countsFor8  int = 0
	countsFor17 int = 0
	frameSpeed  int = 5
	mouseX      int
	mouseY      int
)

//go:embed resource
var asset embed.FS

// GameEngine
func NewGame() *Game {
	bag := storage.New()
	//字体
	font := fonts.NewFont(&asset)
	//场景
	scence := mapManage.NewE1(&asset, bag, font)
	//UI
	ui := layout.NewUI(&asset, font, scence, bag)
	bag.UI = ui
	//BGM
	bgm := music.NewMusicBGM(&asset)

	gameEngine := &Game{
		count:       0,
		countForMap: 0,
		ui:          ui,
		music:       bgm,
		mapManage:   scence,
		font_style:  font,
	}
	//启动游戏
	gameEngine.StartEngine()
	return gameEngine
}

// 引擎启动
func (g *Game) StartEngine() {
	if status.Config.IsNetPlay {
		//网络监听消息
		go g.ListenMessage()
	}
	//隐藏鼠标系统的ICON
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	//UI 初始化
	g.font_style.LoadFont("resource/font/pf_normal.ttf")
	g.ui.LoadGameLoginImages()
	runtime.GC()
}

// 关闭所有连接
func (g *Game) CloseCon() {
	g.Ws.Close()
	close(status.Config.Queue)
}
func (g *Game) Update() error {
	//判断是否是点击屏幕
	if !controller.IsTouch() {
		mouseX, mouseY = ebiten.CursorPosition()
	} else {
		mouseX, mouseY = controller.GetTouchDefaultXY()
	}

	//切换场景状态机
	if !status.Config.ChangeScenceFlg {
		switch status.Config.CurrentGameScence {
		case tools.GAMESCENESTART:
			//进入游戏场景逻辑
			g.changeScenceGameUpdate()
			break
		case tools.GAMESCENEOPENDOOR:
			//游戏加载逻辑
			g.ChangeScenceOpenDoorUpdate()
			break
		case tools.GAMESCENESELECTROLE:
			//进入游戏选择界面逻辑
			g.ChangeScenceSelectUpdate()
			break
		default:
			//进入游戏登录界面逻辑
			g.ChangeScenceLoginUpdate()
			break
		}
	}
	//全屏显示控制
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		i := ebiten.IsFullscreen()
		ebiten.SetFullscreen(!i)
	}
	//手机pc模式切换
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		status.Config.IsMobile = !status.Config.IsMobile
	}
	//Debug 信息显示控制
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		status.Config.DisPlayDebugInfo = !status.Config.DisPlayDebugInfo
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//判断是否切换场景
	if !status.Config.ChangeScenceFlg {
		switch status.Config.CurrentGameScence {
		case tools.GAMESCENESTART:
			g.ChangeScenceGameDraw(screen)
		case tools.GAMESCENESELECTROLE:
			g.ChangeScenceSelectDraw(screen)
		case tools.GAMESCENEOPENDOOR:
			g.ChangeScenceOpenDoorDraw(screen)
		default:
			g.ChangeScenceLoginDraw(screen)
		}
	}

	//PC的场合
	//if !status.Config.IsMobile {
	//绘制鼠标ICON
	g.ui.DrawMouseIcon(screen, mouseX, mouseY)
	//}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tools.LAYOUTX, tools.LAYOUTY
}
