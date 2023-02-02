package main

import (
	"game/engine"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  int = 490
	SCREENHEIGHT int = 300
)

func main() {
	if strings.Compare(runtime.Version(), "go1.19.5") >= 0 {
		debug.SetMemoryLimit(300 * 1024 * 1024)
	}
	//设置log级别
	log.SetFlags(log.Llongfile)
	//设置渲染方法
	os.Setenv("EBITEN_GRAPHICS_LIBRARY", "opengl")
	//配置窗体显示属性
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//size
	ebiten.SetWindowSize(SCREENWIDTH*2, SCREENHEIGHT*2)
	//title
	ebiten.SetWindowTitle("Golang_DibaloⅡ")
	//TPS
	ebiten.SetTPS(60)
	gameStart := engine.NewGame()
	if err := ebiten.RunGame(gameStart); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if gameStart.Ws != nil {
			gameStart.CloseCon()
		}
	}()
}
