package main

import (
	"game/engine"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  int = 490
	SCREENHEIGHT int = 300
)

func main() {
	os.Setenv("EBITEN_GRAPHICS_LIBRARY", "opengl")
	ebiten.SetWindowSize(SCREENWIDTH*2, SCREENHEIGHT*2)
	ebiten.SetWindowTitle("Golang_DibaloⅡ")
	ebiten.SetMaxTPS(80)
	gameStart := engine.NewGame()
	defer func() {
		if gameStart.Ws != nil {
			gameStart.CloseCon()
		}
	}()
	if err := ebiten.RunGame(gameStart); err != nil {
		log.Fatal(err)
	}
}
